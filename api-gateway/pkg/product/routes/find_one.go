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
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	// grpc 요청
	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}