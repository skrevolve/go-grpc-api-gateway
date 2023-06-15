package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductInfoId int64 `json:"productInfoId"`
	Quantity      int64 `json:"quantity"`
}

// 주문 생성
func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {

	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 토큰 인증 미들웨어로 부터 받은 데이터
	userInfoId, _ := ctx.Get("userInfoId")
	// grpc 요청
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductInfoId: body.ProductInfoId,
		Quantity: body.Quantity,
		UserInfoId: userInfoId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}