# go-grpc-api-gateway-microservice

## Configuration Description

1. API Gateway
    - [api-gateway](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/api-gateway)

2. gRPC Service
    - [auth-svc](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/auth-svc)
    - [order-svc](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/order-svc)
    - [product-svc](https://github.com/skrevolve/go-grpc-api-gateway-microservice/tree/master/product-svc)

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
# go install google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Grpc Plugin Issue ([#1070](https://github.com/golang/protobuf/issues/1070))

```sh
protoc protoc/route_guide.proto --go_out=plugins=grpc:.
ERR : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

## VSCode: Colud not import Golang package

use go work in project root

```sh
cd go-grpc-api-gateway-microservice
go work init
go work use ./api-gateway ./auth-svc ./order-svc ./product-svc
```

## Make Resource

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```
