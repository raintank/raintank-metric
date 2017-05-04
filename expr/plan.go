package expr

import (
	"errors"
	"fmt"
	"io"

	"github.com/raintank/metrictank/api/models"
)

type Req struct {
	Query string
	From  uint32 // from for this particular pattern
	To    uint32 // to for this particular pattern
}

type Plan struct {
	Reqs          []Req
	exprs         []*expr
	MaxDataPoints uint32
	From          uint32                  // global request scoped from
	To            uint32                  // global request scoped to
	input         map[Req][]models.Series // input data to work with. set via Run()
	// new data generated by processing funcs. useful for two reasons:
	// 1) reuse partial calculations e.g. queries like target=movingAvg(sum(foo), 10)&target=sum(foo)A (TODO)
	// 2) central place to return data back to pool when we're done.
	generated map[Req][]models.Series
}

func (p Plan) Dump(w io.Writer) {
	fmt.Fprintf(w, "Plan:\n")
	fmt.Fprintf(w, "* Exprs:\n")
	for _, e := range p.exprs {
		fmt.Fprintln(w, e.Print(2))
	}
	fmt.Fprintf(w, "* Reqs:\n")
	for _, r := range p.Reqs {
		fmt.Fprintln(w, "   ", r)
	}
	fmt.Fprintf(w, "MaxDataPoints: %d\n", p.MaxDataPoints)
	fmt.Fprintf(w, "From: %d\n", p.From)
	fmt.Fprintf(w, "To: %d\n", p.To)
}

// Plan validates the expressions and comes up with the initial (potentially non-optimal) execution plan
// which is just a list of requests and the expressions.
// traverse tree and as we go down:
// * make sure function exists
// * tentative validation pre function call (number of args and type of args, to the extent it can be done in advance),
// * let function validate input arguments further (to the extend it can be done in advance)
// * allow functions to extend the notion of which data is required
// * future version: allow functions to mark safe to pre-aggregate using consolidateBy or not
func NewPlan(exprs []*expr, from, to, mdp uint32, stable bool, reqs []Req) (Plan, error) {
	var err error
	for _, e := range exprs {
		reqs, err = newplan(e, from, to, stable, reqs)
		if err != nil {
			return Plan{}, err
		}
	}
	return Plan{
		Reqs:          reqs,
		exprs:         exprs,
		MaxDataPoints: mdp,
		From:          from,
		To:            to,
	}, nil
}

// consumeArg verifies that the argument at pos j matches the expected argType
// it's up to the caller to assure that j is valid before calling.
// if argType allows for multiple arguments, j is advanced to cover all accepted arguments.
// the returned j is always the index where the next argument should be.
func consumeArg(args []*expr, j int, exp argType) (int, error) {
	got := args[j]
	switch exp {
	case series:
		if got.etype != etName && got.etype != etFunc {
			return 0, ErrBadArgumentStr{"func or name", string(got.etype)}
		}
	case seriesList:
		if got.etype != etName && got.etype != etFunc {
			return 0, ErrBadArgumentStr{"func or name", string(got.etype)}
		}
	case seriesLists:
		if got.etype != etName && got.etype != etFunc {
			return 0, ErrBadArgumentStr{"func or name", string(got.etype)}
		}
		// special case! consume all subsequent args (if any) in args that will also yield a seriesList
		for len(args) > j+1 && (args[j+1].etype == etName || args[j+1].etype == etFunc) {
			j += 1
		}
	case integer:
		if got.etype != etInt {
			return 0, ErrBadArgumentStr{"int", string(got.etype)}
		}
	case integers:
		if got.etype != etInt {
			return 0, ErrBadArgumentStr{"int", string(got.etype)}
		}
		// special case! consume all subsequent args (if any) in args that will also yield an integer
		for len(args) > j+1 && args[j+1].etype == etInt {
			j += 1
		}
	case float:
		if got.etype != etFloat && got.etype != etInt {
			return 0, ErrBadArgumentStr{"float", string(got.etype)}
		}
	case str:
		if got.etype != etString {
			return 0, ErrBadArgumentStr{"string", string(got.etype)}
		}
	case boolean:
		if got.etype != etBool {
			return 0, ErrBadArgumentStr{"string", string(got.etype)}
		}
	}
	j += 1
	return j, nil
}

// consumeKwarg consumes the kwarg (by key k) and verifies it
// it's the callers responsability that k exists within namedArgs
// it also makes sure the kwarg has not been consumed already via the kwargs map
// (it would be an error to provide an argument twice via the same keyword,
// or once positionally and once via keyword)
func consumeKwarg(optArgs []optArg, namedArgs map[string]*expr, k string, seenKwargs map[string]struct{}) error {
	var found bool
	var exp optArg
	for _, exp = range optArgs {
		if exp.key == k {
			found = true
			break
		}
	}
	if !found {
		return ErrUnknownKwarg{k}
	}
	_, ok := seenKwargs[k]
	if ok {
		return ErrKwargSpecifiedTwice{k}
	}
	seenKwargs[k] = struct{}{}
	got := namedArgs[k]
	switch exp.val {
	case integer:
		if got.etype != etInt {
			return ErrBadKwarg{k, integer, got.etype}
		}
	case float:
		// integer is also a valid float, just happened to have no decimals
		if got.etype != etInt && got.etype != etFloat {
			return ErrBadKwarg{k, float, got.etype}
		}
	case str:
		if got.etype != etString {
			return ErrBadKwarg{k, str, got.etype}
		}
	}
	return nil
}

// newplan adds requests as needed for the given expr, resolving function calls as needed
func newplan(e *expr, from, to uint32, stable bool, reqs []Req) ([]Req, error) {
	if e.etype != etFunc && e.etype != etName {
		return nil, errors.New("request must be a function call or metric pattern")
	}
	if e.etype == etName {
		reqs = append(reqs, Req{
			e.str,
			from,
			to,
		})
		return reqs, nil
	}

	// here e.type is guaranteed to be etFunc
	fdef, ok := funcs[e.str]
	if !ok {
		return nil, ErrUnknownFunction(e.str)
	}
	if stable && !fdef.stable {
		return nil, ErrUnknownFunction(e.str)
	}

	fn := fdef.constr()
	return newplanFunc(e, fn, from, to, stable, reqs)
}

// newplanFunc adds requests as needed for the given expr, and validates the function input
// provided you already know the expression is a function call to the given function
func newplanFunc(e *expr, fn Func, from, to uint32, stable bool, reqs []Req) ([]Req, error) {
	// first comes the interesting task of validating the arguments as specified by the function,
	// against the arguments that were parsed.

	argsExp, argsOptExp, _ := fn.Signature()
	var err error

	// note:
	// * signature may have seriesLists in it, which means one or more args of type seriesList
	//   so it's legal to have more e.args than signature args in that case.
	// * we can't do extensive, accurate validation of the type here because what the output from a function we depend on
	//   might be dynamically typed. e.g. movingAvg returns 1..N series depending on how many it got as input

	// first validate the mandatory args

	j := 0 // pos in args of next given arg to process
	for _, argExp := range argsExp {
		if len(e.args) <= j {
			return nil, ErrMissingArg
		}
		j, err = consumeArg(e.args, j, argExp)
		if err != nil {
			return nil, err
		}
	}

	// we stopped iterating the mandatory args.
	// any remaining args should be due to optional args otherwise there's too many
	// we also track here which keywords can also be used for the given optional args
	// so that those args should not be specified via their keys anymore.

	seenKwargs := make(map[string]struct{})
	for _, argOpt := range argsOptExp {
		if len(e.args) <= j {
			break // no more args specified. we're done.
		}
		j, err = consumeArg(e.args, j, argOpt.val)
		if err != nil {
			return nil, err
		}
		seenKwargs[argOpt.key] = struct{}{}
	}
	if len(e.args) > j {
		return nil, ErrTooManyArg
	}

	// for any provided keyword args, verify that they are what the function stipulated
	// and that they have not already been specified via their position
	for k := range e.namedArgs {
		err = consumeKwarg(argsOptExp, e.namedArgs, k, seenKwargs)
		if err != nil {
			return nil, err
		}

	}
	err = fn.Init(e.args, e.namedArgs)
	if err != nil {
		return nil, err
	}
	from, to = fn.NeedRange(from, to)
	// look at which arguments are requested
	// if the args are series, they are to be requested with the potentially extended to/from
	// if they are not, keep traversing the tree until we find out which metrics to fetch and for which durations
	for _, arg := range e.args {
		if arg.etype == etName || arg.etype == etFunc {
			reqs, err = newplan(arg, from, to, stable, reqs)
			if err != nil {
				return nil, err
			}
		}
	}
	return reqs, nil
}

// Run invokes all processing as specified in the plan (expressions, from/to) with the input as input
func (p Plan) Run(input map[Req][]models.Series) ([]models.Series, error) {
	var out []models.Series
	p.input = input
	p.generated = make(map[Req][]models.Series)
	for _, expr := range p.exprs {
		o, err := p.run(p.From, p.To, expr)
		if err != nil {
			return nil, err
		}
		out = append(out, o...)
	}
	return out, nil
}

func (p Plan) run(from, to uint32, e *expr) ([]models.Series, error) {
	if e.etype != etFunc && e.etype != etName {
		panic("this should never happen. request must be a function call or metric pattern")
	}
	if e.etype == etName {
		req := Req{
			e.str,
			from,
			to,
		}
		return p.input[req], nil
	}

	// here e.type is guaranteed to be etFunc
	fdef, ok := funcs[e.str]
	if !ok {
		panic(fmt.Sprintf("cannot find func %q. this should never happen as we should have validated function existence earlier", e.str))
	}
	fn := fdef.constr()
	err := fn.Init(e.args, e.namedArgs)
	if err != nil {
		return nil, err
	}
	from, to = fn.NeedRange(from, to)
	// look at which arguments are requested
	// if the args are series, they are to be requested with the potentially extended to/from
	// if they are not, keep traversing the tree until we find out which metrics to fetch and for which durations
	results := make([]interface{}, len(e.args))
	for i, arg := range e.args {
		if arg.etype == etName || arg.etype == etFunc {
			result, err := p.run(from, to, arg)
			if err != nil {
				return nil, err
			}
			results[i] = result
		} else if arg.etype == etString {
			results[i] = arg.str
		} else if arg.etype == etInt {
			results[i] = arg.int
		} else {
			// etype == etFloat
			results[i] = arg.float
		}
	}
	named := make(map[string]interface{})
	for k, arg := range e.namedArgs {
		if arg.etype == etString {
			named[k] = arg.str
		} else if arg.etype == etInt {
			named[k] = arg.int
		} else if arg.etype == etFloat {
			named[k] = arg.float
		} else {
			panic(fmt.Sprintf("named arg cannot be of type %q", arg.etype))
		}
	}

	// we now have all our args and can process the data and return
	rets, err := fn.Exec(p.generated, named, results...)
	if err != nil {
		return nil, err
	}
	series := make([]models.Series, len(rets))
	for i, ret := range rets {
		series[i] = ret.(models.Series)
	}
	return series, nil
}

// Clean returns all buffers (all input data + generated series along the way)
// back to the pool.
func (p Plan) Clean() {
	for _, series := range p.input {
		for _, serie := range series {
			pointSlicePool.Put(serie.Datapoints[:0])
		}
	}
	for _, series := range p.generated {
		for _, serie := range series {
			pointSlicePool.Put(serie.Datapoints[:0])
		}
	}
}
