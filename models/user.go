package models

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserInfoId       uint      	`gorm:"primaryKey"`
	ImgPath          string
	Name             string
	Gender           string
	Social           string
	Email            string
	Phone            string
	Password         string
	Country          string
	Lang             string
	IpAddr           string
	LoginDate        time.Time
	LogoutDate       time.Time
	InsertDate       time.Time
	UpdateDate       time.Time
	Block            bool
}

func (u *User) GenerateHashPassword() error {
	if len(u.Password) == 0 {
		return errors.New("invalid param. password should not be empty")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

func MakeSessionKey(n int) string {
	const (
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}