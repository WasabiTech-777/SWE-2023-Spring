package routes

import (
	"net/http"
)



func GetArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var article models.Article
	initialize.DB.First(&article, params["aid"])

	if article.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Article does not exist"))
	}
	json.NewEncoder(writer).Encode(article)
}

func GetBody(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var article models.Article
	initialize.DB.First(&article, params["aid"])

	if article.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Article does not exist"))
	}
	json.NewEncoder(writer).Encode(article.Url)
}


func PostArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var article models.Article
	json.NewDecoder(request.Body).Decode(&article)
	json.NewEncoder(writer).Encode(article)
}

func PutArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var article models.Article
	initialize.DB.First(&article, params["aid"])
	json.NewDecoder(request.Body).Decode(&article)
	initialize.DB.Save(&article)
	json.NewEncoder(writer).Encode(article)
}

func DeleteArticle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var article models.Article
	initialize.DB.First(&article, params["sid"])

	if article.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Article does not exist"))
	}
	initialize.DB.Unscoped().Delete(&article)
	writer.WriteHeader(http.StatusOK)
}


