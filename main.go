package main

/*-----------TODO-----------*/
//1. Create Rest API
//2. Create a PostgreSQL database with Docker
//3. Enable Sign-In, saving name data and encrypting and saving password data

/*---------RESOURCES--------*/

//Sign-In and Password Implementation: https://www.sohamkamani.com/golang/password-authentication-and-storage/





import (
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"net/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
)


var database *sql.DB

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func main(){
	router := mux.NewRouter()

	//first argument is address of webpage
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("Hello World")
		})

	
	// Handlers in credentials.go
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	
	// initialize database connection
	Init()
	
	// start the server on port 8000
	log.Println("Wi-key!")
	log.Fatal(http.ListenAndServe(":8000", nil))

	
	//http.ListenAndServe(":4000", router)

}

func Init() *gorm.DB {
	
	dsn := "host=localhost user=pg password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"//TODO Database fails to connect, connection refused
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//db.AutoMigrate()

	return database
}

func Signup(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
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

func Signin(w http.ResponseWriter, r *http.Request){
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