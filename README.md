# grpc
gRPC server - client example

protoc --go_out=pkg \
    --go-grpc_out=require_unimplemented_servers=false:pkg \
    api/proto/adder.proto
