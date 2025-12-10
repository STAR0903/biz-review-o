GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
    # 将 where git 找到的 git.exe 路径，直接替换成 bash.exe 的路径
    # Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))  # <--- 原始代码

    # 新的、正确的逻辑：将 `git.exe` 直接替换为 `bash.exe`，并转换斜杠
    # (shell where git) -> E:\command\Git\bin\git.exe
    # (subst git.exe,bash.exe,...) -> E:\command\Git\bin\bash.exe
    # (subst \,/,...) -> E:/command/Git/bin/bash.exe
    Git_Bash=$(subst \,/,$(subst git.exe,bash.exe,$(shell where git)))

    INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
    API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
    INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
    API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: validate
# generate validate proto
validate:
	protoc --proto_path=. \
    		--proto_path=./third_party \
			--go_out=paths=source_relative:. \
        	--validate_out=paths=source_relative,lang=go:. \
           $(API_PROTO_FILES)

.PHONY: errors
errors:
	protoc --proto_path=. \
         --proto_path=./third_party \
         --go_out=paths=source_relative:. \
         --go-errors_out=paths=source_relative:. \
         $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# wire
.PHONY: wire
wire:
	cd cmd/review-o && wire


