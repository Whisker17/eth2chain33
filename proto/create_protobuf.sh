#!/bin/sh

#chain33_path=$(go list -f '{{.Dir}}' "github.com/33cn/chain33")
protoc --go_out=plugins=grpc:../types ./*.proto  --proto_path=.
