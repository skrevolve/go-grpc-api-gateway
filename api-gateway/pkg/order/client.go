package order

import (
	"fmt"

	"github.com/skrevolve/api-gateway/pkg/config"
	"github.com/skrevolve/api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

// order microservice 연결
func InitServiceClient(c *config.Config) pb.OrderServiceClient {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(c.OrderSvcUrl, opts) //no SSL running. grpc.WithInsecure() not supported now
	if err != nil {
		fmt.Println("Colud not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}