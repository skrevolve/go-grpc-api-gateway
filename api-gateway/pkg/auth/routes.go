package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/auth/routes"
	"github.com/skrevolve/api-gateway/pkg/config"
)

// auth 라우팅
func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	// auth microservice 연결
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// http 바인딩
	auth := r.Group("/auth")
	auth.POST("/register", svc.Register)
	auth.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}