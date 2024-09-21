package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/db_golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(MYSQL), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("sucsessfully connected to database")
}