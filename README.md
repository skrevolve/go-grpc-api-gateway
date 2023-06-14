# go-grpc-api-gateway

https://grpc-ecosystem.github.io/grpc-gateway/

## Installing protobuffer

### Linux (Ubuntu)
```
sudo apt install -y protobuf-compiler
```

### MacOS
```
brew install protobuff
```

### Windows
```
goto https://github.com/protocolbuffers/protobuf/releases/tag/v23.1
intsall protoc-23.1-win32.zip or protoc-23.1-win64.zip
```

### GRPC and Protobuffer package dependencies

```sh
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

// module github.com/skrevolve/grpc-gateway