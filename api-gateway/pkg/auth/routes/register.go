package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/skrevolve/api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 회원 가입
func Register(ctx *gin.Context, c pb.AuthServiceClient) {

	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// auth service 서버로 요청
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:		body.Email,
		Password:	body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}