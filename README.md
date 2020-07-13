# grpc-example


[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/bvwells/grpc-example?tab=overview)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bvwells/grpc-example)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bvwells/grpc-example) 
[![Build Status](https://travis-ci.com/bvwells/grpc-example.svg?branch=master)](https://travis-ci.org/bvwells/grpc-example)
[![codecov](https://codecov.io/gh/bvwells/grpc-example/branch/master/graph/badge.svg)](https://codecov.io/gh/bvwells/grpc-example)
[![Go Report Card](https://goreportcard.com/badge/github.com/bvwells/grpc-example)](https://goreportcard.com/report/github.com/bvwells/grpc-example)

This repository contains an example usage of gRPC based on the grpc quickstart example defined [here](https://grpc.io/docs/languages/go/quickstart/) and [here](https://github.com/grpc/grpc-go/blob/master/examples/helloworld/helloworld/helloworld.proto).

## Install protoc

Install protoc (https://github.com/protocolbuffers/protobuf)

```
brew install protobuf
```

## Install tools

To install the go protoc plugin run:

```
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

Ensure that $GOBIN is added to the path.

```
export PATH=$PATH:~/go/bin
```

## Generate client and server stubs

```bash
$ go generate ./...
```

## Run server

To run the server:

```bash
$ cd server
$ go run main.go
```

## Run client

To run the server:

```bash
$ cd client
$ go run main.go
```
