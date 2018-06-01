#!/bin/bash

# This script auto-generates protobuf related files. It is intended to be run manually when either
# API types are added/modified, or server gRPC calls are added. The generated files should then
# be checked into source control.

set -x
set -o errexit
set -o nounset
set -o pipefail

PROJECT_ROOT=$(cd $(dirname ${BASH_SOURCE})/..; pwd)
PATH="${PROJECT_ROOT}/dist:${PATH}"

go build -i -o dist/protoc-gen-go ./vendor/github.com/golang/protobuf/protoc-gen-go
GOPROTOBINARY=go

PROTO_FILES=$(find $PROJECT_ROOT \( -name "*.proto" -and -path '*/pkg/*' -or -name "*.proto" -and -path '*/holidays/*' -or -path '*/calendar/*' -and -name "*.proto" \))
for i in ${PROTO_FILES}; do
    protoc \
        -I${PROJECT_ROOT} \
        -I/usr/local/include \
        -I./vendor \
        -I$GOPATH/src \
        --${GOPROTOBINARY}_out=plugins=grpc:$GOPATH/src \
        $i
done