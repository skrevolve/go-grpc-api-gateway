package middlewares

import (
	"grpc/util"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserInfoId uint
	Email      string
	SessionKey string
	Iat        int
	Exp        int
}

type AppleJwtPayload struct {
	Iss            string
	Aud            string
	Exp            int
	Iat            int
	Sub            string
	Nonce          string
	CHash          string
	Email          string
	EmailVerified  string
	IsPrivateEmail string
	AuthTime       int
	NonceSupported bool
}

var SECRET = os.Getenv("JWT_SECRET_KEY")

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func GenerateAccessToken(userInfoId uint, email, sessionKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_info_id": userInfoId,
		"email": email,
		"session_key": sessionKey,
		"exp": time.Now().Add(time.Hour * 2).Unix(), // 2H
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := accessToken.SignedString([]byte(SECRET))
	if err != nil { return "", util.HttpError("GEN_ACCESS_TOKEN_ERR") }
	return token, nil
}

func GenerateRefreshToken(userInfoId uint, email, sessionKey string) (fiber.Map, error) {
	claims := jwt.MapClaims{
		"user_info_id": userInfoId,
		"email": email,
		"session_key": sessionKey,
		"exp": time.Now().Add(time.Hour * 336).Unix(), // 14D
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := accessToken.SignedString([]byte(SECRET))
	if err != nil { return nil, util.HttpError("GEN_REFRESH_TOKEN_ERR") }
	return fiber.Map{"refresh_token": token}, nil
}