// Code generated by "stringer -type=Format"; DO NOT EDIT.

package msg

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FormatMetricDataArrayJson-0]
	_ = x[FormatMetricDataArrayMsgp-1]
	_ = x[FormatMetricPoint-2]
	_ = x[FormatMetricPointWithoutOrg-3]
	_ = x[FormatIndexControlMessage-4]
}

const _Format_name = "FormatMetricDataArrayJsonFormatMetricDataArrayMsgpFormatMetricPointFormatMetricPointWithoutOrgFormatIndexControlMessage"

var _Format_index = [...]uint8{0, 25, 50, 67, 94, 119}

func (i Format) String() string {
	if i >= Format(len(_Format_index)-1) {
		return "Format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}
