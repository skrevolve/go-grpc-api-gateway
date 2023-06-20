# api gateway

## resource

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## module install

```sh
go mod init github.com/YOUR_USERNAME/api-gateway
go get google.golang.org/grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
```

## Generate proto

```sh
make proto
```

## Run Server

```sh
make server
```