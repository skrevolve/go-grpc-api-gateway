package models

type GetUserInfoByEmailAndPassword struct {
	UserInfoId uint
	Email      string
	Social     string
	Lang 			 string
	Block 	   bool
}