package main

/*-----------TODO-----------*/
//1. Create Rest API
//2. Create a PostgreSQL database with Docker
//3. Enable Sign-In, saving name data and encrypting and saving password data

/*---------RESOURCES--------*/
/*--------------------IMPORTING GORM POSTGRESQL DRIVERS--------------------
documentation here: https://gorm.io/docs/connecting_to_the_database.html

//Sign-In and Password Implementation: https://www.sohamkamani.com/golang/password-authentication-and-storage/
*/


import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
*/
