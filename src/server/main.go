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
	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Authentication", "content-type", os.Getenv("angular_domain")})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// Test hello world
	router.HandleFunc("/", routes.GetHome).Methods("GET")

	// Routes for Session entity
	router.HandleFunc("/session/{sid}", routes.GetSession).Methods("GET")
	router.HandleFunc("/session/user/{uid}", routes.GetSessionFromUser).Methods("GET")
	router.HandleFunc("/session/article/{aid}", routes.GetSessionFromArticle).Methods("GET")
	router.HandleFunc("/session/", routes.PostSession).Methods("OPTIONS", "POST")
	router.HandleFunc("/session/{sid}", routes.PutSession).Methods("PUT")
	router.HandleFunc("/session/{sid}", routes.DeleteSession).Methods("DELETE")

	// Routes for Article entity
	router.HandleFunc("/article/{aid}", routes.GetArticle).Methods("GET")
	router.HandleFunc("/article/", routes.PostArticle).Methods("OPTIONS", "POST")
	router.HandleFunc("/article/{aid}", routes.PutArticle).Methods("PUT")
	router.HandleFunc("/article/{aid}", routes.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/article/body/{aid}", routes.GetBody).Methods("GET")

	// Routes for User entity
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{uid}", routes.GetUser).Methods("GET")
	router.HandleFunc("/uname/{uname}", routes.GetUserFromName).Methods("GET")
	router.HandleFunc("/users", routes.PostUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/users/{uid}", routes.PutUser).Methods("PUT")
	router.HandleFunc("/users/{uid}", routes.DeleteUser).Methods("DELETE")
	router.HandleFunc("/login", routes.AuthenticateUser).Methods("OPTIONS", "POST", "GET")
	router.HandleFunc("/token", routes.ValidateToken).Methods("OPTIONS", "POST", "GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router)))

}
