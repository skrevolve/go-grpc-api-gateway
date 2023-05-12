package models

type GetUserInfo struct {
	UserInfoId uint
	Email      string
	Social     string
	Lang 			 string
	Block 	   bool
}

type LoginResponse struct {
	AccessToken string  `json:"access_token"`
}