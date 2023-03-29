package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/src/server/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello")
	initialize.LoadEnv()
	initialize.Connect()
	initialize.Migrate()
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Credentials", "Authentication", "content-type", os.Getenv("angular_domain")})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	//Test hello world
	router.HandleFunc("/", routes.GetHome).Methods("GET")

	//Routes for User entity
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{uid}", routes.GetUser).Methods("GET")
	router.HandleFunc("/uname/{uname}", routes.GetUserFromName).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/users/{uid}", routes.PutUser).Methods("PUT")
	router.HandleFunc("/users/{uid}", routes.DeleteUser).Methods("DELETE")
	router.HandleFunc("/login", routes.AuthenticateUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/token", routes.ValidateToken).Methods("OPTIONS", "POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins, handlers.AllowCredentials())(router)))

}
