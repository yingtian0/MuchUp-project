PROTO_DIR := api-schema/proto
GO_OUT := gen/go
GOOGLEAPIS := third_party/googleapis
DESCRIPTOR_OUT := api-gateway/envoy/proto/api.pb
TOOL_BIN := $(CURDIR)/.tools/bin
BUF := $(TOOL_BIN)/buf
BUF_CACHE_DIR ?= /tmp/buf-cache
BUF_LLM_TEMPLATE := buf.gen.llm.yaml
BUF_VERSION := v1.59.0
PROTOC_GEN_GO_VERSION := v1.36.11
PROTOC_GEN_GO_GRPC_VERSION := v1.5.1
SWAG_VERSION := v1.16.4

export PATH := $(TOOL_BIN):$(PATH)

PROTO_FILES := \
	$(PROTO_DIR)/chat/v1/chat.proto \
	$(PROTO_DIR)/auth/v1/auth.proto \
	$(PROTO_DIR)/llm/v1/llm.proto

LLM_PROTO_FILE := $(PROTO_DIR)/llm/v1/llm.proto

PROTO_FILE ?=

.PHONY: tools-install proto gen gen-file gen-llm descriptor lint breaking clean

tools-install:
	mkdir -p $(TOOL_BIN)
	GOBIN=$(TOOL_BIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	GOBIN=$(TOOL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	GOBIN=$(TOOL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	GOBIN=$(TOOL_BIN) go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

proto: gen descriptor

gen:
	BUF_CACHE_DIR=$(BUF_CACHE_DIR) $(BUF) generate

gen-file:
	test -n "$(PROTO_FILE)"
	BUF_CACHE_DIR=$(BUF_CACHE_DIR) $(BUF) generate "$(PROTO_FILE)"

gen-llm:
	BUF_CACHE_DIR=$(BUF_CACHE_DIR) $(BUF) generate "$(LLM_PROTO_FILE)" --template $(BUF_LLM_TEMPLATE)

descriptor:
	mkdir -p $(dir $(DESCRIPTOR_OUT))
	protoc \
		-I $(PROTO_DIR) \
		-I $(GOOGLEAPIS) \
		--include_imports \
		--include_source_info \
		--descriptor_set_out=$(DESCRIPTOR_OUT) \
		$(PROTO_FILES)

lint:
	BUF_CACHE_DIR=$(BUF_CACHE_DIR) $(BUF) lint

breaking:
	BUF_CACHE_DIR=$(BUF_CACHE_DIR) $(BUF) breaking --against '.git#branch=main'

clean:
	rm -rf $(GO_OUT)/*
	rm -f $(DESCRIPTOR_OUT)
