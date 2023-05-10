# grpc-go
- go install : https://go.dev/doc/install
- proto3 install : https://grpc.io/docs/protoc-installation/
- docs : https://grpc.io/docs/languages/go/quickstart/

# Installing protobuffer

### Linux

```
sudo apt install -y protobuf-compiler
```

### MacOS

```
brew install protobuff
```

### GRPC and Protobuffer package dependencies

go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

NOTE: You should add the `protoc-gen-go-grpc` to your PATH

```
PATH="${PATH}:${HOME}/go/bin"

```

### Running the service

```
make run
```

### test
>
  http://localhost:3000/?ticker=ETH
>
  http://localhost:3000/?ticker=BTC
>
  http://localhost:3000/?ticker=HANDY
>
  http://localhost:3000/?ticker=TEST