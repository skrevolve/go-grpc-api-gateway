package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"grpc/middlewares"
	"grpc/services"
	"grpc/util"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(*fiber.Ctx, string, string, string) (string, error)
}

type userService struct{}

// is the business logic
func (s *userService) Login(c *fiber.Ctx, email string, password string, device_token string) (string, error) {

	email = strings.ToLower(strings.Trim(email, " "))

	userAgent := util.UserAgentProvider(c)
	if len(userAgent) == 0 { return "", fmt.Errorf("잘못된 디바이스 정보 요청입니다") }

	validPassword := util.ValidPassword(password)
	if !validPassword { return "", fmt.Errorf("비밀번호 형식이 잘못되었습니다") }

	hasher := sha256.New()
	hasher.Write([]byte(password))
	password = hex.EncodeToString(hasher.Sum(nil))

	userInfo, err := services.GetUserInfo(email, password)
	if err != nil { log.Fatal(err) }
	if len(userInfo) == 0 { return "", fmt.Errorf("아이디 또는 비밀번호가 잘못 되었습니다") }
	if userInfo[0].Block { return "", fmt.Errorf("차단된 계정입니다") }
	if userInfo[0].Social != "normal" { return "", fmt.Errorf("이미 등록된 이메일 입니다") }

	sessionKey := util.MakeRamdomString(13)
	access_token, _ := middlewares.GenerateAccessToken(userInfo[0].UserInfoId, userInfo[0].Email, sessionKey)

	return access_token, nil
}

type loggingService struct {
	next UserService
}

func (s loggingService) Login(c *fiber.Ctx, email string, password string, device_token string) (access_token string, err error) {

	// defer func(begin time.Time) {
	// 	logrus.WithFields(logrus.Fields{
	// 		"email": ,
	// 		"took":      time.Since(begin),
	// 		"err":       err,
	// 		"price":     price,
	// 		"ticker":    ticker,
	// 	}).Info("FetchPrice")
	// }(time.Now())

	return s.next.Login(c, email, password, device_token)
}