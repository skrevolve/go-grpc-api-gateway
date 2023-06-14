proto:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.

# protoc -I ./proto \
# 	 --go_out ./proto --go_opt paths=source_relative \
#    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
#    --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \ # https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
#    ./proto/helloworld/hello_world.proto

server:
	go run ./cmd/main.go
