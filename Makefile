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
