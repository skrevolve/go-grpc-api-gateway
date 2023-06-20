build:
	go build -o bin/micro

run: build
	bin/micro

proto:
	protoc pkg/pb/*.proto --go_out=. --go-grpc_out=.

# https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/

# protoc -I pkg/pb \
# 	 --go_out pkg/pb --go_opt paths=source_relative \
#    --go-grpc_out pkg/pb --go-grpc_opt paths=source_relative \
#    --grpc-gateway_out pkg/pb --grpc-gateway_opt paths=source_relative \
#    pkg/pb/*.proto

server:
	go run cmd/main.go

.PHONY: proto