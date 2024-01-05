#!/bin/bash

# Check if buf is installed
if ! command -v buf &> /dev/null
then
    echo "buf could not be found. Please install it and try again."
    exit 1
fi

# Directory containing your .proto files
PROTO_DIR="api/proto/"

# Lint the .proto files
buf lint "${PROTO_DIR}"
