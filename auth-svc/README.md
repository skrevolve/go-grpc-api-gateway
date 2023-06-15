# api gateway

## 리소스 생성

```sql
create database auth_svc;
create database order_svc;
create database product_svc;
```

## 설치

```sh
go mod init github.com/YOUR_USERNAME/auth-svc
go get github.com/spf13/viper
go get google.golang.org/grpc
go get gorm.io/gorm
go get gorm.io/driver/postgres
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