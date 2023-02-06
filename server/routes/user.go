package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WasabiTech-777/SWE-2023-Spring/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/models"
	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []models.User
	initialize.DB.Find(&users)
	json.NewEncoder(writer).Encode(&users)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user models.User
	initialize.DB.First(&user, params["uid"])

	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User does not exist"))
	}

	json.NewEncoder(writer).Encode(user)
}

func PostUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	post := initialize.DB.Create(&user)
	err := post.Error

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	}

	json.NewEncoder(writer).Encode(user)
}

func PutUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user models.User
	initialize.DB.First(&user, params["uid"])
	json.NewDecoder(request.Body).Decode(&user)
	initialize.DB.Save(&user)
	json.NewEncoder(writer).Encode(user)

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user models.User
	initialize.DB.First(&user, params["uid"])

	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User does not exist"))
	}
	initialize.DB.Unscoped().Delete(&user)
	writer.WriteHeader(http.StatusOK)
}
