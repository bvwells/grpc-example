package main

//go:generate protoc -I. --go_out=plugins=grpc,paths=source_relative:./ helloworld/hello-world.proto
