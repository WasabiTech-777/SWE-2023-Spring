package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, err = DB.Query("insert into users values ($1, $2)", user.Username, string(hashedPassword)); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `User` instance
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the existing entry present in the database for the given username
	result := DB.Find("select password from users where username=$1", user.Username)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We create another instance of `User` to store the user we get from the database
	storedUser := &User{}
	// Store the obtained password in `storedUser`
	if result := DB.Find("this hashed password=$1").First(&user.Password); result.Error != nil {
		// error handling...
	}
	err = result.Scan(&storedUser.Password) //FIXME
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		err := DB.First(&user, 100).Error

		if errors.Is(err, gorm.ErrRecordNotFound) { //FIXME
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}
