package database

import (
	"log"
	"time"

	"github.com/FSDelite/go-rest-api/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := "root:senha123@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error: ", err)
	}
	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
