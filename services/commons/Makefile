# make gen-proto src=user
# make gen-proto src=auth
# generate protos/user.proto protobuf to protos/userpb
.PHONY: gen-proto
gen-proto:
ifeq ($(src),)
	@echo "Error: Please specify the proto source using src=<name>"
	@exit 1
endif
	protoc --go_out=./protos \
	--go-grpc_out=./protos \
	--proto_path=./protos $(src).proto
