package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

)

type JwtCredentials struct {
	Name string `json:"uname"`
	jwt.StandardClaims
}

var jwtkey = []byte(os.Getenv("JWT_key"))




func GetHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "API Home")
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []models.User
	initialize.DB.Find(&users)
	json.NewEncoder(writer).Encode(&users)
}

func AuthenticateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var userCredentials models.User
	json.NewDecoder(request.Body).Decode(&userCredentials)

	var storedUser models.User
	initialize.DB.Where("name = ?", userCredentials.Name).First(&storedUser)
	//result := initialize.DB.Find("this hashed password=$1").First(&user.Pass)
	//result.Scan(&storedUser.Pass)

	if storedUser.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Pass), []byte(userCredentials.Pass))

	if err != nil {
		// If the two passwords don't match, return a 401 status
		writer.WriteHeader(http.StatusUnauthorized)
	} else {
		//fmt.Fprintf(writer, "hello! you have been authenticated")

		// Settign expiration time of jwt token for one day
		expiration := time.Now().Add(time.Hour * 24)

		// Declare custom claims using the username 
		claims := JwtCredentials{
			Name: storedUser.Name,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiration.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString(jwtkey)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Println(signedToken)
			http.SetCookie(writer, &http.Cookie{
			Name: "token",
			Value: signedToken,
			Expires: expiration,
		})
		}

	}
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

	//json.NewDecoder(request.Body).Decode(&user)
	//AuthenticateUser(&user, writer) //Should this be a function or part of GetUser?

	json.NewEncoder(writer).Encode(user)

}

func GenerateHashedPassword(user *models.User) {
	pwCost := 16
	//original password strings cannot exceed 72 bytes
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), pwCost)
	if err != nil {
		fmt.Println("Password is too long! Maximum length is 72 characters.")
	}

	mismatch := bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Pass))
	if mismatch != nil {
		fmt.Println("Hash failed!")
	}

	user.Pass = string(hashedPassword)
}

func PostUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	GenerateHashedPassword(&user)
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
