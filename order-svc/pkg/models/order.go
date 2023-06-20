package models

type Order struct {
	OrderInfoId   int64 `json:"order_info_id" gorm:"primaryKey"`
	Price         int64 `json:"price"`
	ProductInfoId int64 `json:"product_info_id"`
	UserInfoId    int64 `json:"user_info_id"`
}