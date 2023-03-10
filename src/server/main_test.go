package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/routes"
)

func TestLoadEnv(test *testing.T) {
	initialize.LoadEnv()
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")
	if dsn == "" || port == "" {
		test.Errorf("FAILED. DSN and PORT not sourced")
	} else {
		test.Logf("PASSED. DSN and PORT sourced")
	}
}

func TestConnect(test *testing.T) {
	initialize.LoadEnv()
	db := initialize.Connect()
	ch, _ := db.DB()
	err := ch.Ping()
	if err != nil {
		test.Errorf(`FAILED. %q`, err)
	} else {
		test.Logf("PASSED. Database connection successful")
	}
}

func TestMigrate(test *testing.T) {
	initialize.LoadEnv()
	db := initialize.Connect()
	initialize.Migrate()
	var tableNames = [4]string{"articles", "friends", "sessions", "users"}
	for _, value := range tableNames {
		if !(db.Migrator().HasTable(value)) {
			test.Errorf("Migrate() FAILED. Table %q was not created", value)
		}
	}
	test.Logf("Migrate() PASSED. All tables were migrated")

}

func TestGET(test *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(routes.GetHome))
	defer ts.Close()
	response, err := http.Get(ts.URL)
	if err != nil {
		test.Error(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		test.Error(err)
	}

	message := string(body)

	if message == "API Home" {
		test.Logf("PASSED. Basic GET request done on base route")
	} else {
		test.Errorf("FAILED. Basic Get request on base route, message mis-match")
	}
}

func TestPOST(test *testing.T) {
	var user models.User
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "POST" {
			test.Errorf("Expected method POST, got %s", r.Method)
		}

		routes.PostUser(w, r)

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			test.Error(err)
		}

		// Check that the request body is valid JSON
		err = json.Unmarshal(body, &user)
		if err != nil {
			test.Error(err)
		}

		// Check that the user fields are correct
		if user.Name != "Gator1" {
			test.Errorf("Expected name Gator1, got %s", user.Name)
		}

		// Write a response with a status code of 201 Created
		w.WriteHeader(http.StatusCreated)

	}))
	defer ts.Close()

	// Make a POST request to the test server
	reqBody := []byte(`{"ID": 1000, "uname": "Gator1", "pass": "Gator"}`)
	resp, err := http.Post("http://localhost:9000/users", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != 200 {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPUT(test *testing.T) {
	var user models.User
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "POST" {
			test.Errorf("Expected method POST, got %s", r.Method)
		}

		routes.PutUser(w, r)

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			test.Error(err)
		}

		// Check that the request body is valid JSON
		err = json.Unmarshal(body, &user)
		if err != nil {
			test.Error(err)
		}

		// Check that the user fields are correct
		if user.Name != "Gator2" {
			test.Errorf("Expected name Gator2, got %s", user.Name)
		}

		// Write a response with a status code of 201 Created
		w.WriteHeader(http.StatusOK)

	}))
	defer ts.Close()

	// Make a PUT request to the test server
	reqBody := []byte(`{"ID": 1000, "uname": "Gator2", "pass": "Gator"}`)
	req, err := http.NewRequest("PUT", "http://localhost:9000/users/100000", bytes.NewBuffer(reqBody))

	if err != nil {
		test.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != 200 {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestAuthenicateUser(test *testing.T) {
	//Creating a new server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			test.Error(err)
		}

		// Check that the request body is valid JSON, and create new user
		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			test.Error(err)
		}
		routes.AuthenticateUser(w, r)
	}))
	defer ts.Close()
	//Post request to create new user
	reqBody := []byte(`{"uname": "Gator2", "pass": "Gator"}`)
	resp, err := http.Post("http://localhost:9000/users", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		test.Error(err)
	}

	//If new user was made successfully. authenticate that user with valid credentials
	if resp.StatusCode == 0 {
		req, err := http.NewRequest("AuthenticateUser", "http://localhost:9000/login", bytes.NewBuffer(reqBody))
		if err != nil {
			test.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		// Check that the response status code is 200
		if req.Response.StatusCode != 200 {
			test.Errorf("handler returned wrong status code: got %v want %v",
				req.Response.StatusCode, http.StatusOK)
		}
	}
}

func TestDELETE(test *testing.T) {
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "DELETE" {
			test.Errorf("Expected method DELETE, got %s", r.Method)
		}

		routes.DeleteUser(w, r)

		// Write a response with a status code of 201 Created
		w.WriteHeader(http.StatusNoContent)

	}))
	defer ts.Close()

	// Make a POST request to the test server
	reqBody := []byte(`{}`)
	req, err := http.NewRequest("DELETE", "http://localhost:9000/users/1000", bytes.NewBuffer(reqBody))

	if err != nil {
		test.Error(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != 200 {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPutOld(test *testing.T) {
	var user models.User
	// Create a new test server with a handler function that handles the PUT request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is PUT
		if r.Method != "PUT" {
			test.Errorf("Expected method PUT, got %s", r.Method)
		}

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			test.Error(err)
		}

		// Check that the request body is valid JSON
		err = json.Unmarshal(body, &user)
		if err != nil {
			test.Error(err)
		}

		// Check that the user fields are correct
		if user.Name != "Albert" {
			test.Errorf("Expected name Albert, got %s", user.Name)
		}

		// Write a response with a status code of 200 OK
		w.WriteHeader(http.StatusOK)
	}))

	// Editing the new user
	putReqBody := []byte(`{"uname": "Albert", "pass": "Gator"}`)
	req, err := http.NewRequest("PUT", ts.URL+"/users/"+fmt.Sprint(100000), bytes.NewBuffer(putReqBody))
	if err != nil {
		test.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestDeleteOld(test *testing.T) {
	// Create a new test server with a handler function that handles the DELETE request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is DELETE
		if r.Method != "DELETE" {
			test.Errorf("Expected method DELETE, got %s", r.Method)
		}

		// Check that the request URL is correct
		if r.URL.Path != "/users/1" {
			test.Errorf("Expected URL /users/1, got %s", r.URL.Path)
		}

		// Write a response with a status code of 204 No Content
		w.WriteHeader(http.StatusNoContent)
	}))

	// Make a DELETE request to the test server
	req, err := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	if err != nil {
		test.Error(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 204 No Content
	if resp.StatusCode != http.StatusNoContent {
		test.Errorf("Expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}
