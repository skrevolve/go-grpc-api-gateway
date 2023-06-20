package models

type User struct {
	UserInfoId int64  `json:"user_info_id" gorm:"primaryKey"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}