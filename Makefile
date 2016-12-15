.PHONY: test bin docker
default:
	$(MAKE) all
test:
	CGO_ENABLED=1 go test -v -race $(shell go list ./... | grep -v /vendor/)
check:
	$(MAKE) test
bin:
	./scripts/build.sh
docker:
	./scripts/build_docker.sh
all:
	$(MAKE) bin
	$(MAKE) docker
