#!/bin/sh

readonly go_gen_path="../server/users/gen/proto/"
readonly curr_dir=$(pwd)
readonly dart_gen_path="${curr_dir}/../client/lib/gen/proto/"
readonly proto_file_path="${curr_dir}/auth.proto"

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

protoc --go_out=$go_gen_path --go_opt=paths=source_relative \
    --go-grpc_out=$go_gen_path --go-grpc_opt=paths=source_relative \
    auth.proto

dart pub global activate protoc_plugin

export PATH="$PATH:$HOME/.pub-cache/bin"

protoc --proto_path="$(pwd)" --dart_out=grpc:"$dart_gen_path" "$proto_file_path"
