package main

/*-----------TODO-----------*/
//1. Create Rest API
//2. Create a PostgreSQL database with Docker
//3. Enable Sign-In, saving name data and encrypting and saving password data

/*---------RESOURCES--------*/

//Sign-In and Password Implementation: https://www.sohamkamani.com/golang/password-authentication-and-storage/
//GO Library for handling password creation parameters: https://pkg.go.dev/github.com/sethvargo/go-password/password

import (
	//"database/sql"
	//"encoding/json"
	"log"
	"net/http"

	//"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"golang.org/x/crypto/bcrypt"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

const PORT = ":5432"

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(PORT, router))

}

func main() {

	InitialMigration()
	initializeRouter()

	//first argument is address of webpage
	//router.HandleFunc("/wikey", func(w http.ResponseWriter, r *http.Request) {
	//	json.NewEncoder(w).Encode("Hello World")}

	//http.HandleFunc("/signin", Signin)
	//http.HandleFunc("/signup", Signup)

	// initialize database connection
	//Init()

	// start the server on port 8000
	log.Println("Wi-key!")

	//http.ListenAndServe(":4000", router)
}

/*
var database *sql.DB

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}
*/

/*
func Init() *gorm.DB {

	dsn := "host=localhost user=pg password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai" //TODO Database fails to connect, connection refused
	//"host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(Credentials{}) //if Credentials is in a different file, its &file/packagename.Credentials

	return database
}

func Signup(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Users{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, err = database.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the existing entry present in the database for the given username
	result := database.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}
*/
/*
var DB *gorm.DB
var err error
const DNS = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

type User struct {
	gorm.Model
	FirstName 	string `json:"firstname"`
	LastName 	string `json:"lastname"`
	Email 		string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database!")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, router *http.Request) {

}

func GetUser(w http.ResponseWriter, router *http.Request) {

}

func CreateUser(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(router.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, router *http.Request) {

}

func DeleteUser(w http.ResponseWriter, router *http.Request) {

}
*/
const hello = "Hello"
