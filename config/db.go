package config

import (
	"app/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	// database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&models.Product{})
	db = database
}

func GetDB() *gorm.DB {
	return db
}
