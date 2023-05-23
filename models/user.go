package models


type User struct {
	UserInfoId       uint      `gorm:"primaryKey;not null"`
	Username         string
	Email            string    `gorm:"uniqueKey"`
	Password         string
	Bio              string
	Image            string
}