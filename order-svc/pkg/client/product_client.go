package client

import (
	"context"
	"fmt"

	"github.com/skrevolve/order-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(url, opts) //no SSL running. grpc.WithInsecure() not supported now
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient {
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) FindOne(productInfoId int64) (*pb.FindOneResponse, error) {

	req := &pb.FindOneRequest{
		ProductInfoId: productInfoId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(productInfoId int64, orderInfoId int64) (*pb.DecreaseStockResponse, error) {

    req := &pb.DecreaseStockRequest{
        ProductInfoId: productInfoId,
        OrderInfoId: orderInfoId,
    }

    return c.Client.DecreaseStock(context.Background(), req)
}