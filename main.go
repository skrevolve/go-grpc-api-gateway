package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/skrevolve/grpc/protos/user"
)
type GRPCServer struct {
	userpb.UnimplementedUserServiceServer
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) Login(ctx context.Context, in *userpb.LoginRequest) (*userpb.LoginReply, error) {
	return &userpb.LoginReply{AccessToken: in.Email + "_token"}, nil
}

func (s *GRPCServer) GetProfile(ctx context.Context, in *userpb.GetProfileRequset) (*userpb.GetProfileReply, error) {
	return &userpb.GetProfileReply{Email: "sukyu0919@naver.com"}, nil
}

func main() {
	// listening ont TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	// create a gRPC server object
	s := grpc.NewServer()
	// attach the User service to the server
	userpb.RegisterUserServiceServer(s, &GRPCServer{})
	// serve gRPC Server
	log.Println("serving gRPC on 0.0.0.0:8080")

	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// register UserService
	err = userpb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr: ":8090",
		Handler: gwmux,
	}

	log.Println("serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}