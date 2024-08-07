# Makefile for Your Golang Monorepo Project
PROJECT_NAME := $(shell basename $(CURDIR))

# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

VERSION := $(shell git describe --tags --always)

# Targets
.PHONY: all help version
.PHONY: lint clean

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

version: ## show version
	@echo $(VERSION)

.PHONY: dev
dev: ## run dev server
	docker compose up --build

lint: ## run golangci-lint
	@golangci-lint run ./...

clean: ## clean build directory
	@rm -rf cover.out result.json ./deployments/charts/*.tgz
	@rm -rf $(BUILD_DIR)

.PHONY: build
build: ## build go binary
	@go build -v ./...

.PHONY: test
test: ## run test
	@go test -v ./...

.PHONY: coverage
coverage: ## generate coverage
	@go test -json -coverprofile=cover.out ./... >result.json

.PHONY: gen-pb
gen-pb: ## generate protobuf
	@echo Starting generate pb
	@find ./pb -name '*.proto' | while read -r proto_file; do \
		protoc --proto_path=./pb \
		    --go_out=paths=source_relative:./ \
		    --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:./ \
		    --go-grpc-mock_out=paths=source_relative,require_unimplemented_servers=false:./ \
		    $$proto_file; \
	done
	@echo Successfully generated proto

	@echo Starting inject tags
	@protoc-go-inject-tag -input="./entity/domain/*/model/*.pb.go"
	@echo Successfully injected tags

.PHONY: gen-swagger
gen-swagger: ## generate swagger
	@#swag init -q -g impl.go -d ./adapter/block/scan,./pkg -o ./api/block/scan --instanceName block_scan

### testing
.PHONY: test-api-order
test-api-order: ## test api
	@#k6 run --vus=1 --iterations=1 ./tests/k6/order.api.test.js

.PHONY: test-api-user
test-api-user: ## test api user
	@#k6 run --vus=1 --iterations=1 ./tests/k6/user.api.test.js

.PHONY: test-stress
test-stress: ## test load
	@#k6 run --env SCENARIO=peak_load ./tests/k6/order.api.test.js --out=cloud

.PHONY: test-load
test-load: ## test stress
	@#k6 run --env SCENARIO=average_load ./tests/k6/order.api.test.js --out=cloud

## docker
IMAGE_NAME := ghcr.io/blackhorseya/$(PROJECT_NAME)

.PHONY: docker-push
docker-push: ## push docker image
	@echo "Pushing Docker image to $(IMAGE_NAME):$(VERSION)"
	docker buildx build --push \
  --tag $(IMAGE_NAME):latest \
  --tag $(IMAGE_NAME):$(VERSION) .

## deployments
DEPLOY_TO := prod
HELM_REPO_NAME := blackhorseya

#.PHONY: deploy
#deploy: deploy-app deploy-storage ## deploy all
#
#.PHONY: deploy-app
#deploy-app: deploy-restaurant-restful ## deploy app
#
#.PHONY: deploy-restaurant-restful
#deploy-restaurant-restful: ## deploy restaurant
#	@helm upgrade $(DEPLOY_TO)-godine-restaurant-restful $(HELM_REPO_NAME)/godine \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-restaurant-restful.yaml
#
#.PHONY: deploy-storage
#deploy-storage: deploy-mariadb deploy-mongodb deploy-redis ## deploy storage
#
#.PHONY: deploy-mariadb
#deploy-mariadb: ## deploy mariadb
#	@helm upgrade $(DEPLOY_TO)-godine-mariadb bitnami/mariadb \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-mariadb.yaml
#
#.PHONY: deploy-mongodb
#deploy-mongodb: ## deploy mongodb
#	@helm upgrade $(DEPLOY_TO)-godine-mongodb bitnami/mongodb \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-mongodb.yaml
#
#.PHONY: deploy-redis
#deploy-redis: ## deploy redis
#	@helm upgrade $(DEPLOY_TO)-godine-redis bitnami/redis \
#  --install --namespace $(PROJECT_NAME) \
#  --history-max 3 \
#  --values ./deployments/$(DEPLOY_TO)/godine-redis.yaml
