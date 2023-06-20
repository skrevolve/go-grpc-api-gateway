package models

type StockDecreaseLog struct {
	Id           int64 `json:"id" gorm:"primaryKey"`
	OrderInfoId  int64 `json:"order_info_id"`
	ProductRefer int64 `json:"product_info_id"`
}