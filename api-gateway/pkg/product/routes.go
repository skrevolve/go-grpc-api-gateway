package product

import (
	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/auth"
	"github.com/skrevolve/api-gateway/pkg/config"
	"github.com/skrevolve/api-gateway/pkg/product/routes"
)

// product 라우팅
func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {

	// auth 미들웨어 연결
	a := auth.InitAuthMiddleware(authSvc)

	// product microservice 연결
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// http 바인딩
	routes := r.Group("/product")
	routes.Use(a.AuthRequired) // 토큰 인증 미들웨어 사용
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:product_info_id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}