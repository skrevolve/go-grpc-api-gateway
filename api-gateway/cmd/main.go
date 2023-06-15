package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/auth"
	"github.com/skrevolve/api-gateway/pkg/config"
	"github.com/skrevolve/api-gateway/pkg/order"
	"github.com/skrevolve/api-gateway/pkg/product"
)

func main() {

	c, err := config.LoadConfig() // 환경 설정 실행
	if err != nil {
		log.Fatalln("Failed at config loading", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}