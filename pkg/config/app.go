package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "bartek:bxra2@13mg@tcp(localhost)/simple_rest?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the MySQL database
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB { return db }
