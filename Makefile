## env for project
PROJECT_NAME := $(shell basename $(PWD))

## common
.PHONY: check-%
check-%: ## check environment variable is exists
	@if [ -z '${${*}}' ]; then echo 'Environment variable $* not set' && exit 1; fi

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:  ## remove artifacts
	@rm -rf coverage.txt profile.out ./bin ./deployments/charts/*.tgz
	@echo Successfuly removed artifacts

## go
.PHONY: gazelle-repos
gazelle-repos: ## update gazelle repos
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -prune

.PHONY: gazelle
gazelle: gazelle-repos ## run gazelle with bazel
	@bazel run //:gazelle

.PHONY: test-go
test-go: ## bazel test
	@bazel test //...

.PHONY: build-go
build-go: ## bazel build
	@bazel build //...

## generate
.PHONY: gen-pb
gen-pb: ## generate protobuf messages and services
	@go get -u google.golang.org/protobuf/proto
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go

	## Starting generate pb
	@protoc --proto_path=. \
			--go_out=. --go_opt=module=github.com/blackhorseya/ryze \
			--go-grpc_out=. --go-grpc_opt=module=github.com/blackhorseya/ekko,require_unimplemented_servers=false \
			./pb/domain/*/**.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./pkg/entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-mock
gen-mock: ## generate mock
	@go generate ./...
	## Successfully generated mock
