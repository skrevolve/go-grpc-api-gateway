package auth

import (
	"fmt"

	pb "github.com/skrevolve/api-gateway/pkg/auth/pb"
	"github.com/skrevolve/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

// auth microservice 연결 실행
func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(c.AuthSvcUrl, opts) //no SSL running. grpc.WithInsecure() not supported now
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}