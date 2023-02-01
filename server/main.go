package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

var DB *gorm.DB

func main() {
	LoadEnv()
	InitializeConnection()
	router := mux.NewRouter()

	//Test hello world
	router.HandleFunc("/", HelloHandler).Methods("GET")

	//Routes for User entity
	//router.HandleFunc("/users", GetUserHandler).Methods("GET")
	router.HandleFunc("/users", PostUserHandler).Methods("POST")

	//outer.HandleFunc("/users/{uid}", GetUserHandler).Methods("GET")
	//router.HandleFunc("/users/{uid}", PutUserHandler).Methods("PUT")
	//router.HandleFunc("/users/{uid}", DeleteUserHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(os.Getenv(":9000"), router))

}
