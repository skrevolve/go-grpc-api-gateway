package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Mysql *gorm.DB
}

var Conn DBInstance

func Init() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	schema := os.Getenv("MYSQL_SCHEMA")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	charset := os.Getenv("MYSQL_CHARSET")

	var err error
	var config gorm.Config
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, schema, charset,
	)

	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		config = gorm.Config{}
	} else {
		config = gorm.Config{ Logger: logger.Default.LogMode(logger.Silent) }
	}

	db, err := gorm.Open(mysql.Open(dsn), &config)

	if err != nil {
		log.Fatal("failed to connect database \n", err)
		os.Exit(2)
	}

	Conn = DBInstance{
		Mysql: db,
	}
}