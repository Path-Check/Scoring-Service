#!/bin/bash

protoc --proto_path=pb --go_out=plugins=grpc:pb pb/notification.proto

# protoc --proto_path=pb --go_out=plugins=grpc:pb pb/log.proto

# grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../pb/js/ --grpc_out=../pb/js --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` notification.proto