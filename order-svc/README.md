# api gateway

## 리소스 생성

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## 설치

```sh
go mod init github.com/YOUR_USERNAME/order-svc
go get google.golang.org/grpc
#go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt
```

## Generate proto

```sh
make proto
```

## Run Server

```sh
make server
```