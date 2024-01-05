#!/bin/bash

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc could not be found. Please install it and try again."
    exit 1
fi

# Check if protoc-gen-lint is installed
if ! command -v protoc-gen-lint &> /dev/null
then
    echo "protoc-gen-lint could not be found. Please install it with the following command:"
    echo "GO111MODULE=on go get github.com/ckaznocha/protoc-gen-lint"
    echo "Then, make sure that the directory containing the protoc-gen-lint executable is in your PATH."
    echo "If you installed protoc-gen-lint using the go get command, the executable is likely in your Go bin directory, which is usually $HOME/go/bin or $GOPATH/bin."
    echo "You can add this directory to your PATH with the following command:"
    echo "export PATH=$PATH:$(go env GOPATH)/bin"
    exit 1
fi

# Directory containing your .proto files
PROTO_DIR="api/proto/v1"

# For each .proto file in the directory
for PROTO_FILE in "${PROTO_DIR}"/*.proto
do
  # Lint the .proto file
  protoc -I ${PROTO_DIR} --lint_out=. "${PROTO_FILE}"
done