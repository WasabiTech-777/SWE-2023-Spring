package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "postgres://postgres:postgres@localhost:5432/db"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	//Email     string `json:"email"` Maybe later??
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database!")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, router *http.Request) {
	fmt.Println("meow")
}

func GetUser(w http.ResponseWriter, router *http.Request) {
	fmt.Println("meow")
}

func CreateUser(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(router.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, router *http.Request) {
	fmt.Println("meow")
}

func DeleteUser(w http.ResponseWriter, router *http.Request) {
	fmt.Println("meow")
}
