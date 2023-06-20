package db

import (
	"log"
	"time"

	"github.com/skrevolve/product-svc/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.StockDecreaseLog{})

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return Handler{db}
}