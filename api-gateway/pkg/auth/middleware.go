package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	pb "github.com/skrevolve/api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

// auth 미들웨어 연결
func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

// auth 토큰 인증
func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userInfoId", res.UserInfoId)
	ctx.Next()
}
