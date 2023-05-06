## env for project
PROJECT_NAME := $(shell basename $(PWD))
VERSION := $(shell git describe --tags --always)

## env for helm
HELM_REPO_NAME := sean-side

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
	@bazel test //... --define=VERSION=$(VERSION)

.PHONY: build-go
build-go: ## bazel build
	@bazel build //... --define=VERSION=$(VERSION)

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

## helm
.PHONY: lint-helm
lint-helm: ## lint helm chart
	@helm lint deployments/charts/*

.PHONY: add-helm-repo
add-helm-repo: ## add helm repo
	@helm repo add --no-update $(HELM_REPO_NAME) gs://sean-helm-charts/charts
	@helm repo update $(HELM_REPO_NAME)

.PHONY: package-helm
package-helm: ## package helm chart
	@helm package ./deployments/charts/$(PROJECT_NAME) --destination ./deployments/charts

.PHONY: push-helm
push-helm: ## push helm chart to gcs
	@helm gcs push --force ./deployments/charts/$(PROJECT_NAME)-*.tgz $(HELM_REPO_NAME)
	@helm repo update $(HELM_REPO_NAME)

## docker
.PHONY: push-ryze-restful-image
push-ryze-restful-image: ## push ryze restful image to gcr
	@echo "Starting push ryze restful image version: $(VERSION)"
	@bazel run //:$@ --define=VERSION=$(VERSION)

.PHONY: push-ryze-listener-block-image
push-ryze-listener-block-image: ## push ryze restful image to gcr
	@echo "Starting push ryze restful image version: $(VERSION)"
	@bazel run //:$@ --define=VERSION=$(VERSION)
