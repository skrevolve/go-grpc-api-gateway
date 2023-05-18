package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/skrevolve/grpc/database"
	userpb "github.com/skrevolve/grpc/protos/user"
)

var (
	tcpPort  = flag.String("tcp", ":8080", "listen address of the tcp transport")
	grpcPort = flag.String("grpc", ":8090", "listen address of the gRPC transport")
	ctx      = context.Background()
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
	// connect MySQL
	database.Init()

	// listening ont TCP port
	lis, err := net.Listen("tcp", *tcpPort)
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
		ctx,
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	// register UserService
	err = userpb.RegisterUserServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr: *grpcPort,
		Handler: gwmux,
	}

	log.Println("serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}