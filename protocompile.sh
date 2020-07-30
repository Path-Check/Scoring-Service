#!/bin/bash

protoc --proto_path=pb --go_out=plugins=grpc:pb pb/log.proto --experimental_allow_proto3_optional
protoc --proto_path=pb --go_out=plugins=grpc:pb pb/notification.proto

