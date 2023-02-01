package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitializeConnection() *gorm.DB {
	var err error
	dsn := os.Getenv("DSN")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to the database")
	} else {
		fmt.Println("Success connecting to database!")
	}
	return DB
}

func Migrate() {
	var err error
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic("Error migrating users")
	} else {
		fmt.Println("Success!")
	}
	err = DB.AutoMigrate(&Friends{})
	if err != nil {
		panic("Error migrating friends list")
	}

	err = DB.AutoMigrate(&Articles{})
	if err != nil {
		panic("Error migrating articles")
	}

	err = DB.AutoMigrate(&TypingSessions{})
	if err != nil {
		panic("Error migrating typing sessions")
	}
}
