package database

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/skrevolve/grpc-gateway/models"
)

func dsn() (string, error) {
	envErr := godotenv.Load()
	if envErr != nil {
		return "Error loading .env file", envErr
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("$DB_HOST is not set")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.New("$DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return "", errors.New("$DB_PASSWORD is not set")
	}

	schema := os.Getenv("DB_SCHEMA")
	if schema == "" {
		return "", errors.New("$DB_SCHEMA is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("$DB_PORT is not set")
	}

	charset := os.Getenv("DB_CHARSET")
	if charset == "" {
		return "", errors.New("$DB_CHARSET is not set")
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, schema, charset), nil
}

func New() (*gorm.DB, error) {
	s, err := dsn()
	if err != nil {
		return nil, err
	}

	var db *gorm.DB
	db, err = gorm.Open("mysql", s)
	// defer db.DB().Close()
	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 2)
	// 		db.DB().Ping()
	// 	}
	// }()

	db.DB().SetMaxIdleConns(3)
	db.DB().SetMaxOpenConns(4)
	// show global variables like 'wait_timeout'; => 28800(8 hours)
	db.DB().SetConnMaxLifetime(time.Hour) // 커넥션 유지 최대 유지 시간을 1시간으로 설정
	db.LogMode(false)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
	).Error
	if err != nil {
		return err
	}
	return nil
}