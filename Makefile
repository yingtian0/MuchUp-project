PROTO_DIR=api-schema/proto
GO_OUT=gen/go
GOOGLEAPIS=third_party/googleapis


.PHONY: gen lint breaking clean proto

gen:
	buf generate

lint:
	buf lint

breaking:
	buf breaking --against '.git#branch=main'

clean:
	rm -rf gen/*

proto:
	protoc \
	  -I $(PROTO_DIR) \
	  -I $(GOOGLEAPIS) \
	  --go_out $(GO_OUT) --go_opt paths=source_relative \
	  --go-grpc_out $(GO_OUT) --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out $(GO_OUT) --grpc-gateway_opt paths=source_relative \
	  $(PROTO_DIR)/chat/v1/chat.proto \
	  $(PROTO_DIR)/auth/v1/auth.proto

