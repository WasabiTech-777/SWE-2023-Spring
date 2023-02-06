package main

/*-----------TODO-----------*/
//1. Create Rest API
//2. Create a PostgreSQL database with Docker
//3. Enable Sign-In, saving name data and encrypting and saving password data

/*---------RESOURCES--------*/

//Sign-In and Password Implementation: https://www.sohamkamani.com/golang/password-authentication-and-storage/
//GO Library for handling password creation parameters: https://pkg.go.dev/github.com/sethvargo/go-password/password

import (
	//"database/sql"
	//"encoding/json"
	"log"
	"net/http"

	//"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"golang.org/x/crypto/bcrypt"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

const PORT = ":5432"

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(PORT, router))

}

func main() {

	InitialMigration()
	initializeRouter()

	//first argument is address of webpage
	//router.HandleFunc("/wikey", func(w http.ResponseWriter, r *http.Request) {
	//	json.NewEncoder(w).Encode("Hello World")}

	//http.HandleFunc("/signin", Signin)
	//http.HandleFunc("/signup", Signup)

	// initialize database connection
	//Init()

	// start the server on port 8000
	log.Println("Wi-key!")

	//http.ListenAndServe(":4000", router)
}
