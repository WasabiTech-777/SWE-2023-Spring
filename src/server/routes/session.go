package routes

import (
	"encoding/json"
	"net/http"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/gorilla/mux"
)

func GetSession(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var session models.Session
	initialize.DB.First(&session, params["sid"])

	if session.SessionID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Session does not exist"))
	}
	json.NewEncoder(writer).Encode(session)
}

func GetSessionFromUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var session models.Session
	initialize.DB.Where("SessionID = ?", params["sid"]).Find(&session)
	json.NewEncoder(writer).Encode(session)
}

func GetSessionFromArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var session models.Session
	initialize.DB.Where("ArticleIDD = ?", params["aid"]).Find(&session)
	json.NewEncoder(writer).Encode(session)
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
	var session models.Session
	initialize.DB.First(&session, params["sid"])
	json.NewDecoder(request.Body).Decode(&session)
	initialize.DB.Save(&session)
	json.NewEncoder(writer).Encode(session)
}

func DeleteSession(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var session models.Session
	initialize.DB.First(&session, params["sid"])

	if session.SessionID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Session does not exist"))
	}
	initialize.DB.Unscoped().Delete(&session)
	writer.WriteHeader(http.StatusOK)
}