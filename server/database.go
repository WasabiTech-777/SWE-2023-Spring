package services

import (
	"fmt"
	"github.com/WasabiTech-777/SWE-2023-Spring/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitializeConnection() {
	var err error
	dsn := os.Getenv("DSN")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to the database")
	} else {
		fmt.Println("Success connecting to database!")
	}
	DB.AutoMigrate(&models.User{})
}

func Migrate(DB *gorm.DB) {
	var err error
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error migrating users")
	} else {
		fmt.Println("Success!")
	}
	err = DB.AutoMigrate(&models.Friends{})
	if err != nil {
		panic("Error migrating friends list")
	}

	err = DB.AutoMigrate(&models.Articles{})
	if err != nil {
		panic("Error migrating articles")
	}

	err = DB.AutoMigrate(&models.TypingSessions{})
	if err != nil {
		panic("Error migrating typing sessions")
	}
}
