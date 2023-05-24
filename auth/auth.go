package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

var (
	JWT_ACCESS_SECRET_KEY  = os.Getenv("JWT_ACCESS_SECRET_KEY")
	JWT_REFRESH_SECRET_KEY = os.Getenv("JWT_REFRESH_SECRET_KEY")
)

type claims struct {
	UserInfoId uint 		`json:"user_info_id"`
	Email      string   `json:"email"`
	SessionKey string   `json:"session_key"`
	jwt.StandardClaims
}

func GenerateAccessToken(id uint, email string, sessionKey string) (string, error) {
	return generateAccessToken(id, email, sessionKey, time.Now())
}

func GenerateRefreshToken(id uint, email string, sessionKey string) (string, error) {
	return generateRefreshToken(id, email, sessionKey, time.Now())
}

func generateAccessToken(id uint, email string, sessionKey string, now time.Time) (string, error) {
	claims := &claims{
		id,
		email,
		sessionKey,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 2).Unix(), // 2H
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(JWT_ACCESS_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func generateRefreshToken(id uint, email string, sessionKey string, now time.Time) (string, error) {
	claims := &claims{
		id,
		email,
		sessionKey,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 336).Unix(), // 14D
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(JWT_REFRESH_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func ParseAccessToken(ctx context.Context) (*claims, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "Token")
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_ACCESS_SECRET_KEY, nil
	})
	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("invalid token: it's not a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("token expired")
			} else {
				return nil, fmt.Errorf("invalid token: couldn't handle this token; %w", err)
			}
		} else {
			return nil, fmt.Errorf("invalid token: couldn't handle this token; %w", err)
		}
	}

	c, ok := token.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid token: cannot map token to claims")
	}

	if c.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return c, nil;
}