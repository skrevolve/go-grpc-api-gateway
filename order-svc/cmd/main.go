package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skrevolve/order-svc/pkg/client"
	"github.com/skrevolve/order-svc/pkg/config"
	"github.com/skrevolve/order-svc/pkg/db"
	"github.com/skrevolve/order-svc/pkg/pb"
	"github.com/skrevolve/order-svc/pkg/services"
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
        log.Fatalln("Failed to listing:", err)
    }

    productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    fmt.Println("Order Svc on", c.Port)

    s := services.Server{
        H:          h,
        ProductSvc: productSvc,
    }

    grpcServer := grpc.NewServer()

    pb.RegisterOrderServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalln("Failed to serve:", err)
    }
}