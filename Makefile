PROTO_DIR=api-schema/proto
  GO_OUT=gen/go
  GOOGLEAPIS=third_party/googleapis
  DESCRIPTOR_OUT=api-gateway/envoy/proto/api.pb

  PROTO_FILES = \
        $(PROTO_DIR)/chat/v1/chat.proto \
        $(PROTO_DIR)/auth/v1/auth.proto

  .PHONY: proto gen descriptor lint breaking clean

  proto: gen descriptor

  gen:
        buf generate

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
        buf lint

  breaking:
        buf breaking --against '.git#branch=main'

  clean:
        rm -rf $(GO_OUT)/*
        rm -f $(DESCRIPTOR_OUT)