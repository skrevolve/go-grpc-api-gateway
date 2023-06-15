# go-grpc-api-gateway

<!-- https://grpc-ecosystem.github.io/grpc-gateway/ -->

## Installing protobuffer

### Linux (Ubuntu)

```sh
sudo apt install -y protobuf-compiler
```

### MacOS

```sh
brew install protobuff
```

### Windows

```sh
> goto https://github.com/protocolbuffers/protobuf/releases/tag/v23.1
> intsall protoc-23.1-win32.zip or protoc-23.1-win64.zip
```

### GRPC and Protobuffer package dependencies

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## issue ([link click](https://github.com/golang/protobuf/issues/1070))

```sh
protoc protoc/route_guide.proto --go_out=plugins=grpc:.
OUTPUT : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```