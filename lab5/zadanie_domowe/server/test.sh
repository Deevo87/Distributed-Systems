#!/usr/bin/zsh

JAVA_PLUGIN_PATH="/home/rlaskowski/Pulpit/studia/rozprochy/Distributed-Systems/lab5/zadanie_domowe/server/protoc-gen-grpc-java"
JAVA_OUT="src/main/java"

mkdir -p $JAVA_OUT
protoc -I . --plugin=$JAVA_PLUGIN_PATH --java_out=$JAVA_OUT --grpc-java_out=$JAVA_OUT ExecutionService.proto