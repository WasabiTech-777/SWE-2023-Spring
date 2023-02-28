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
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "POST" {
			test.Errorf("Expected method POST, got %s", r.Method)
		}

		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			test.Error(err)
		}

		// Check that the request body is valid JSON
		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			test.Error(err)
		}

		// Check that the user fields are correct
		if user.Name != "Albert" {
			test.Errorf("Expected name Albert, got %s", user.Name)
		}

		// Write a response with a status code of 201 Created
		w.WriteHeader(http.StatusCreated)
	}))

	// Make a POST request to the test server
	reqBody := []byte(`{"uname": "Albert", "pass": "Gator"}`)
	resp, err := http.Post(ts.URL+"/users", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 201 Created
	if resp.StatusCode != http.StatusCreated {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

}

func TestPUT(test *testing.T) {
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
		if user.Name != "Alberta" {
			test.Errorf("Expected name Alberta, got %s", user.Name)
		}

		// Write a response with a status code of 200 OK
		w.WriteHeader(http.StatusOK)
	}))

	// Editing the new user
	userID := user.ID
	putReqBody := []byte(`{"uname": "Alberta", "pass": "Gator"}`)
	req, err := http.NewRequest("PUT", ts.URL+"/users/"+fmt.Sprint(userID), bytes.NewBuffer(putReqBody))
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

func TestDELETE(test *testing.T) {
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

/*
func TestAuthenicateUser(t *testing.T) {
	req := []byte(`{"uname": "Albatross", "pass": "Gator"}`)

	request, err := http.NewRequest("POST", "/path/to/handler", bytes.NewBuffer(req))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(routes.AuthenticateUser)
	handler.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := `{"message":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
*/

func HelperCreateUser(ts *httptest.Server, reqBody []byte, test *testing.T) {
	// Creating the new user
	postReq, err := http.Post(ts.URL+"/users", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		test.Error(err)
	}
	if postReq.StatusCode != http.StatusCreated {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, postReq.StatusCode)
	}
}
