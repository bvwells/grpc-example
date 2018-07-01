# gRPC-fun
[![GoDoc](http://godoc.org/github.com/bvwells/gRPC-fun?status.svg)](http://godoc.org/github.com/bvwells/gRPC-fun)
[![Build Status](https://travis-ci.org/bvwells/gRPC-fun.svg?branch=master)](https://travis-ci.org/bvwells/gRPC-fun)
[![codecov](https://codecov.io/gh/bvwells/gRPC-fun/branch/master/graph/badge.svg)](https://codecov.io/gh/bvwells/gRPC-fun)
[![Go Report Card](https://goreportcard.com/badge/github.com/bvwells/gRPC-fun)](https://goreportcard.com/report/github.com/bvwells/gRPC-fun)

Fun with gRPC

Build Protobuffer API with:

```
docker run -it -v ${PWD}:/go/src/gRPC-fun -w /go/src/gRPC-fun bvwells/golang-proto protoc -I helloworld/ helloworld/hello-world.proto --go_out=plugins=grpc:helloworld
```
