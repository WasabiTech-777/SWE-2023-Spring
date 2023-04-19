package routes

import (
	"encoding/json"
	"net/http"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/gorilla/mux"
)

func GetSession(writer http.ResponseWriter, request *http.Request) {
	//TODO
}


func PostSession(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var session models.Session
	json.NewDecoder(request.Body).Decode(&session)
	json.NewEncoder(writer).Encode(session)
}

func PutSession(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var session models.User
	initialize.DB.First(&session, params["sid"])
	json.NewDecoder(request.Body).Decode(&session)
	initialize.DB.Save(&session)
	json.NewEncoder(writer).Encode(session)
}

func DeleteSession(writer http.ResponseWriter, request *http.Request) {
	//TODO
}

func GetSessionFromUser(writer http.ResponseWriter, request *http.Request) {
	//TODO
}

func GetSessionFromArticle(writer http.ResponseWriter, request *http.Request) {
	//TODO
}