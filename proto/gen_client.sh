#!/bin/sh

readonly proto_file=$1
readonly dart_gen_path=$2

readonly curr_dir=$(pwd)
readonly proto_file_path="${curr_dir}/${proto_file}.proto"

export PATH="$PATH:$HOME/.local/bin"

dart pub global activate protoc_plugin

export PATH="$PATH:$HOME/.pub-cache/bin"

protoc --proto_path="$(pwd)" --dart_out=grpc:"$dart_gen_path" "$proto_file_path"
