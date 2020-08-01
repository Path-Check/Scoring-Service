#!/bin/bash

protoc --proto_path=pb --go_out=plugins=grpc:pb pb/log.proto
protoc --proto_path=pb --go_out=plugins=grpc:pb pb/notification.proto

