#!/bin/bash

# Check if buf is installed
if ! command -v buf &> /dev/null
then
    echo "buf could not be found. Please install it and try again."
    exit 1
fi

# Generate the Go server code
buf generate api/proto

# Navigate to the pkg/idlefantasystory directory
cd pkg/idlefantasystory || exit 1

# Run go mod tidy
go mod tidy
