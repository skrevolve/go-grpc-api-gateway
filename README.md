# go-grpc-api-gateway-microservice

## microservice spec

- go-grpc
- Gin
- mysql
- gorm

## configuration description

- api-gateway: API server (localhost:8080)
- auth-svc: Auth Service server (localhost:50051)
- order-svc: Order Service server (localhost:50052)
- product-svc: Product Service server (localhost:50053)

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
go install google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## issue ([link click](https://github.com/golang/protobuf/issues/1070))

```sh
protoc protoc/route_guide.proto --go_out=plugins=grpc:.
OUTPUT : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

## VSCode: Colud not import Golang package

use go work in project root

```sh
cd go-grpc-api-gateway-microservice
go work init
go work use ./api-gateway ./auth-svc ./order-svc ./product-svc
```
