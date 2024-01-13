# Contributing to idle-fantasy-story

This project is only to have fun and learn while practicing Go and gRPC.
This document provides guidelines and instructions for contributing.

## Build, lint, test, and run
We use `make` to build, lint, test, and run the project and, in general, to automate any recurrent task.
If a task requires something more complex than a single command, we create a script in the `scripts` directory.
You can have a full list of commands by running `make help`.

We use `revive` for linting.

## Generating the mock code
We use [mockery](https://vektra.github.io/mockery/latest) to generate the mock code.
Follow the instructions in the mockery documentation to install it.

## Modifying the gRPC code
The gRPC code is located in the `api/proto` directory. You can modify the `.proto` files in this directory according to your needs.
We use [buf](https://buf.build/docs/introduction) to lint and generate the code from the `.proto` files.

You can lint the proto files by running:
```bash
make lint-proto
```

You can generate the code by running:
```bash
make generate-proto
```
