package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type JwtCredentials struct {
	Name string `json:"uname"`
	jwt.StandardClaims
}

var jwtkey = []byte(os.Getenv("JWT_key"))

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

func AuthenticateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
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
				Name:    "token",
				Value:   signedToken,
				Expires: expiration,
			})
		}

	}
}

func ValidateToken(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	} else {

		fmt.Println("Parsed Token")
		tokenValue := cookie.Value
		claims := &JwtCredentials{}

		token, err := jwt.ParseWithClaims(tokenValue, claims, func(tkn *jwt.Token) (interface{}, error) { return jwtkey, nil })

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				writer.WriteHeader(http.StatusUnauthorized)
			}
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			if token.Valid {
				writer.Header().Set("Content-Type", "application/json")
				json.NewEncoder(writer).Encode(claims)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
			}

		}
	}
}

func GetHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "API Home")
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

	//json.NewDecoder(request.Body).Decode(&user)
	//AuthenticateUser(&user, writer) //Should this be a function or part of GetUser?

	json.NewEncoder(writer).Encode(user)

}

func GetUserFromName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user models.User
	user.Name = params["uname"]
	fmt.Println("Username: " + user.Name)
	//search for user with name in params
	result := initialize.DB.Where("name = ?", params["uname"]).First(&user)
	//result := initialize.DB.First(&user, params["uname"])
	if result.Error != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User does not exist"))
	}
	json.NewEncoder(writer).Encode(user)
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
	GenerateHashedPassword(&user)
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
