#!/bin/sh

readonly proto_file=$1
readonly go_gen_path=$2
readonly dart_gen_path=$3

readonly curr_dir=$(pwd)
readonly proto_file_path="${curr_dir}/${proto_file}.proto"

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

protoc --go_out="$go_gen_path" --go_opt=paths=source_relative \
    --go-grpc_out="$go_gen_path" --go-grpc_opt=paths=source_relative \
    "${proto_file}".proto

dart pub global activate protoc_plugin

export PATH="$PATH:$HOME/.pub-cache/bin"

protoc --proto_path="$(pwd)" --dart_out=grpc:"$dart_gen_path" "$proto_file_path"
