#!/usr/bin/zsh

protoc --proto_path=../server/ --go_out=generated --go-grpc_out=generated ExecutionService.proto