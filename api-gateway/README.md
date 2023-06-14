# api gateway

## 리소스 생성

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## 설치

```sh
go mod init github.com/YOUR_USERNAME/go-grpc-api-gateway
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# google.golang.org/protobuf/cmd/protoc-gen-go -> google.golang.org/protobuf/cmd/protoc-gen-go-grpc 로 변경되었음
# protoc protoc/route_guide.proto --go_out=plugins=grpc:. // 플러그인 지원 안함
# OUTPUT : --go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

issue : <https://github.com/golang/protobuf/issues/1070>