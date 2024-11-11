package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Content string
}

func GetDB() (*gorm.DB, error) {
	connStr := os.Getenv("DATABASE")

	return gorm.Open(postgres.Open(connStr), &gorm.Config{})
}

func InitDB() {
	db, err := GetDB()
	if err != nil {
		fmt.Println(err)
	}

	error := db.AutoMigrate(&Note{})
	if error != nil {
		fmt.Println(err)
	}
}
