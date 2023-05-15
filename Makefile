build:
	go build -o bin/micro

run: build
	./bin/micro

proto:
	protoc -I ./protos \
	--go_out ./protos --go_opt paths=source_relative \
	--go-grpc_out ./protos --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./protos --grpc-gateway_opt paths=source_relative \
	protos/user/user.proto

.PHONY: proto