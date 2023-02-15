package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WasabiTech-777/SWE-2023-Spring/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/routes"
	"github.com/gorilla/mux"
)

func main() {
	initialize.LoadEnv()
	initialize.Connect()
	initialize.Migrate()
	router := mux.NewRouter()
	//Test hello world
	router.HandleFunc("/", routes.HelloHandler).Methods("GET")

	//Routes for User entity
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{uid}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("POST")
	router.HandleFunc("/users/{uid}", routes.PutUser).Methods("PUT")
	router.HandleFunc("/users/{uid}", routes.DeleteUser).Methods("DELETE")
	router.HandleFunc("/login", routes.AuthenticateUser)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
