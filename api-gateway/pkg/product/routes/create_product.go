package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/product/pb"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

// 제품 생성
func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {

	body := CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// product service 서버로 요청
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequset{
		Name: body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}