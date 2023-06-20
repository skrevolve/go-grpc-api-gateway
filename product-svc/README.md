# Product Service

## Package Dependencies

```sh
# go mod init github.com/YOUR_USERNAME/product-svc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/grpc
go get github.com/spf13/viper
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

## Generate proto

```sh
make proto
```

## Run Server [http://localhost:50052]

```sh
make server
```
