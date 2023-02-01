package main

import (
	"encoding/json"
	"fmt"
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
	var user User
	DB.Find(&user)
	json.NewEncoder(writer).Encode(user)
}

func PostUserHandler(writer http.ResponseWriter, router *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(router.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func PutUserHandler(writer http.ResponseWriter, router *http.Request) {
	//TODO
}

func DeleteUserHandler(writer http.ResponseWriter, router *http.Request) {
	//TODO
}
