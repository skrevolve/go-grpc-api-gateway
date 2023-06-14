package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skrevolve/api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 로그인 HTTP 요청 바인딩
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// auth RPC microservice 로그인으로 전달
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}