# api gateway

## 리소스 생성

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## 설치

```sh
# go mod init github.com/YOUR_USERNAME/go-grpc-api-gateway
go get github.com/gin-gonic/gin
go get github.com/spf13/viper
go get google.golang.org/grpc
```

## Generate proto

```sh
make proto
```

## Run Server

```sh
make server
```