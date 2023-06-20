package models

type Product struct {
	ProductInfoId     int64            `json:"procut_info_id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Stock             int64            `json:"stock"`
	Price             int64            `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}