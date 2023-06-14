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
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper
go get -u google.golang.org/grpc
# go get -u google.golang.org/protobuf/cmd/protoc-gen-go // 플러그인 지원 안함. 아래로 변경
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## issue ([link click](https://github.com/golang/protobuf/issues/1070))

```sh
protoc protoc/route_guide.proto --go_out=plugins=grpc:.
OUTPUT : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```