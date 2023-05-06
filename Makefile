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

.PHONY: update-paackage
update-package: ## update package
	@echo Starting update package
	@go get -u ./...
	@go mod tidy

	@echo Starting update bazel dependencies
	$(MAKE) gazelle-repos

	@git add go.mod go.sum deps.bzl
	@git commit -m "build: update package"
	@echo Successfully updated package

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
			--go-grpc_out=. --go-grpc_opt=module=github.com/blackhorseya/ryze,require_unimplemented_servers=false \
			./pb/domain/*/**.proto
	@echo Successfully generated proto

	## Starting inject tags
	@protoc-go-inject-tag -input="./pkg/entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-mock
gen-mock: ## generate mock
	@go generate ./...
	## Successfully generated mock

.PHONY: gen-wire
gen-wire: ## generate wire
	@wire gen ./...

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -q --dir ./cmd/restful,./ -o ./api/docs
	@echo Successfully generated swagger spec
