package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skrevolve/product-svc/pkg/config"
	"github.com/skrevolve/product-svc/pkg/db"
	"github.com/skrevolve/product-svc/pkg/pb"
	"github.com/skrevolve/product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config loading", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}
	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H:	h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}