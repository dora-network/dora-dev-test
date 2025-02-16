# Check for a .env file for environment variables and if one exists, export those
# for use in the Make environment.
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.DEFAULT_GOAL := proto

PROTOC_GEN_GO_VERSION ?= v1.34.1
GRPC_GATEWAY_VERSION ?= v2.20.0
BUF_VERSION ?= v1.32.1

compose-up:
	@docker compose -f build/env/test-network.yml up --build -d

compose-down:
	@docker compose -f build/env/test-network.yml down

install-dependencies:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@${GRPC_GATEWAY_VERSION}
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${GRPC_GATEWAY_VERSION}
	@go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}

buf-clean:
	@find api -name '*.pb.go' | xargs -r rm
	@find docs/openapiv2 -name '*.swagger.json' | xargs -r rm

buf-lint:
	@buf lint

mod-update:
	@buf dep update

buf-breaks:
	@buf breaking --against ./.git#branch=main,subdir=./

buf-push:
	@buf push

buf-build:
	@buf build

.PHONY: proto
proto: install-dependencies buf-clean buf-lint mod-update
	@buf generate

