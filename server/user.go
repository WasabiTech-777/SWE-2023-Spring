package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/WasabiTech-777/SWE-2023-Spring/models"
	"github.com/WasabiTech-777/SWE-2023-Spring/services"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, router *http.Request) {
	fmt.Fprintf(writer, "hello")
}

func GetUserHandler(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	param := mux.Vars(router)
	for _, element := range param {
		fmt.Println(element)
	}
	var users models.User
	services.DB.Find(&users)
	json.NewEncoder(writer).Encode(users)
}

func PostUserHandler(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(router.Body).Decode(&user)
	services.DB.Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func PutUserHandler(writer http.ResponseWriter, router *http.Request) {
	//TODO
}

func DeleteUserHandler(writer http.ResponseWriter, router *http.Request) {
	//TODO
}
