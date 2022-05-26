#!/bin/sh

protoc --proto_path=./codes/protocol --go_out=plugins=grpc:. ./codes/protocol/*.proto

mv github.com/atopx/serverlib/codes/* ./codes/ && rm -rf github.com

