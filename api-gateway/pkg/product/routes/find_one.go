package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/product/pb"
)

// 제품 찾기
func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {

	// param 정수 변환
	productInfoId, _ := strconv.ParseInt(ctx.Param("product_info_id"), 10, 32)

	// product service 서버로 요청
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		ProductInfoId: int64(productInfoId),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}