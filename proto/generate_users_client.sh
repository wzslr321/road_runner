#!/bin/sh

readonly proto_file=$1
readonly dart_gen_path=$2

readonly curr_dir=$(pwd)
readonly proto_file_path="${curr_dir}/${proto_file}.proto"

PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"

dart pub global activate protoc_plugin

export PATH="$PATH:$HOME/.pub-cache/bin"

protoc --proto_path="$(pwd)" --dart_out=grpc:"$dart_gen_path" "$proto_file_path"
