package order

import (
	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/auth"
	"github.com/skrevolve/api-gateway/pkg/config"
	"github.com/skrevolve/api-gateway/pkg/order/routes"
)

// order 라우팅
func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {

	// auth 미들웨어 연결
	a := auth.InitAuthMiddleware(authSvc)

	// order microservice 연결
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// http 바인딩
	routes := r.Group("/order")
	routes.Use(a.AuthRequired) // 토큰 인증 미들웨어 사용
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}