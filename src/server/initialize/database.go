package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/WasabiTech-777/SWE-2023-Spring/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	var err error
	dsn := os.Getenv("DSNonline")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Error connecting to the database")
	} else {
		fmt.Println("Success connecting to database!")
	}
	return DB
}

func Migrate() {
	var err error
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		panic("Error migrating users")
	}
	err = DB.AutoMigrate(&models.Friends{})
	if err != nil {
		panic("Error migrating friends list")
	}

	err = DB.AutoMigrate(&models.Article{})
	if err != nil {
		panic("Error migrating articles")
	}

	err = DB.AutoMigrate(&models.Session{})
	if err != nil {
		panic("Error migrating typing sessions")
	}
}
