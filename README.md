# Helloworld in Go gRPC
If you not familiar with gRPC you may want to read [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction) first.

## Prerequisites
- Go, we were using version 1.15.x when this document was created. For installation instructions, see Go’s [Getting Started](https://golang.org/doc/install) Guide
- [Protocol buffer](https://developers.google.com/protocol-buffers) compiler, `protoc`, [version 3](https://developers.google.com/protocol-buffers/docs/proto3). For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation)
- Go plugin for the protocol compiler:

Install the protocol compiler plugin for Go (protoc-gen-go) using the following command:
```bash
export GO111MODULE=on # Enable module mode
go get github.com/golang/protobuf/protoc-gen-go
```

Update your PATH so that the protoc compiler can find the plugin:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## How to compile generated proto buffer in Go
```bash
protoc --go_out=plugins=grpc:../. protobuf/helloworld.proto
```

generate grpc-gateway
```bash
protoc --grpc-gateway_out ../. protobuf/helloworld.proto
```

generate swagger config
```bash
protoc --swagger_out . protobuf/helloworld.proto
```

or generate everything at once
```bash
protoc --go_out=plugins=grpc:../. --grpc-gateway_out ../. --swagger_out . protobuf/helloworld.proto
```

## How to run
### Run the gRPC server
```bash
go run server/main.go
```

### Run the REST server
```bash
go run rest/main.go
```

### Run the client
```bash
go run client/main.go [name]
```