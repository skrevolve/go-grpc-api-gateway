package models

import (
	"errors"
	"time"

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