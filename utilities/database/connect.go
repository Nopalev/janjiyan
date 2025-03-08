package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var dsn strings.Builder
	var err error
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	fmt.Fprintf(&dsn, "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
