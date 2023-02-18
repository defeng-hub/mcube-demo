PROJECT_NAME := "mcube-demo"
MAIN_FILE_PAHT := "main.go"
PKG := "github.com/defeng-hub/mcube-demo"
IMAGE_PREFIX := "github.com/defeng-hub/mcube-demo"

MOD_DIR := C:/Users/Administrator/go/pkg/mod

MCUBE_MODULE := "github.com/infraboard/mcube"
MCUBE_VERSION :=$(shell go list -m ${MCUBE_MODULE} | cut -d' ' -f2)
MCUBE_PKG_PATH := ${MOD_DIR}/${MCUBE_MODULE}@${MCUBE_VERSION}

BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"

.PHONY: all dep lint vet test test-coverage build clean

all: build

v:
	echo ${MCUBE_PKG_PATH}

tidy: ## Get the dependencies
	@go mod tidy


build: dep ## Build the binary file
	@go build -a -o dist/aaa -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MAIN_FILE}

linux: dep ## Build the binary file
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MAIN_FILE}

init: dep ## Inital project 
	@go run main.go init

run: install dep ## Run Server
	@go run main.go start

clean: ## Remove previous build
	@go clean .
	@rm -f dist/${PROJECT_NAME}

install: ## Install depence go package
	@go install github.com/infraboard/mcube/cmd/mcube@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest

pb: ## Copy mcube protobuf files to common/pb
	@mkdir -pv common/pb/github.com/infraboard/mcube/pb
	@cp -r ${MCUBE_PKG_PATH}/pb/* common/pb/github.com/infraboard/mcube/pb
	@sudo rm -rf common/pb/github.com/infraboard/mcube/pb/*/*.go

gen: ## Init Service
	@protoc -I=. -I=common/pb --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} apps/*/pb/*.proto
	@go fmt ./...

	@protoc-go-inject-tag -input=apps/*/*.pb.go
	@mcube generate enum -p -m apps/*/*.pb.go


help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'