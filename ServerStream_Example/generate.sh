#! bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc --go_out=. --go-grpc_out=. greet/greetpb/greet.proto