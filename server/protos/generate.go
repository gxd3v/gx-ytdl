package protos

//go:generate F:/protoc/bin/protoc.exe --proto_path=../../protos --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ../../protos/codes/actions.proto ../../protos/codes/errors.proto ../../protos/codes/success.proto ../../protos/codes/prefix.proto ../../protos/global.proto ../../protos/messages.proto ../../protos/server.proto
