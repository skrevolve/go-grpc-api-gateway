package product

import (
	"fmt"

	"github.com/skrevolve/api-gateway/pkg/config"
	"github.com/skrevolve/api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(c.ProductSvcUrl, opts) //no SSL running. grpc.WithInsecure() not supported now
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}