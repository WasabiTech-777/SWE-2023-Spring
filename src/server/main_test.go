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
	"time"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/src/server/routes"
	"github.com/gorilla/mux"
)

// Global Variables for Testing
var Cookie string
var TestUserID int = -1

// Global constant for Server Route
const SERVER_ROUTE = "http://localhost:9000"

// *********************************************LOADING TESTS****************************************//
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

//*********************************************routes.user TESTS****************************************//

func TestGetUsers(test *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(routes.GetUsers))
	//Check that the method is GET
	defer ts.Close()

	resp, err := http.Get(SERVER_ROUTE + "/users")

	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

// NOTE: Running any of the tests below individually will not work. Use command "go test" to run all tests at once
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
	reqBody := []byte(`{"uname": "Gator1", "pass": "Gator"}`)
	resp, err := http.Post((SERVER_ROUTE + "/users"), "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var newUser models.User
	// Parse the response body for the new user's ID
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		test.Error(err)
	}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		test.Error(err)
	}

	TestUserID = int(newUser.ID)
}

func TestGetUser(test *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(routes.GetUser))

	defer ts.Close()

	if TestUserID == -1 {
		test.Errorf("FAILED. TestUserID for Test User not set (User does not exist)")
	}

	resp, err := http.Get(SERVER_ROUTE + "/users/" + fmt.Sprintf("%d", TestUserID))

	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestGetUserFromName(test *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(routes.GetUserFromName))
	//Check that the method is GET
	defer ts.Close()

	resp, err := http.Get(SERVER_ROUTE + "/uname/Gator1")

	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != http.StatusOK {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPUT(test *testing.T) {
	var user models.User
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "PUT" {
			test.Errorf("Expected method PUT, got %s", r.Method)
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

		// Write a response with a status code of 200 OK
		w.WriteHeader(http.StatusOK)

	}))
	defer ts.Close()

	// Make a PUT request to the test server
	reqBody := []byte(`{"Name": "Gator2", "Password": "Gator"}`)
	//reqBody := []byte(`{"uname": "Gator2", "pass": "Gator"}`)
	req, err := http.NewRequest("PUT", (SERVER_ROUTE + "/users/" + fmt.Sprintf("%d", TestUserID)), bytes.NewBuffer(reqBody))

	if err != nil {
		test.Error(err)
	}

	//req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		test.Error(err)
	}

	// Check that the response status code is 200 Success
	if resp.StatusCode != 200 {
		test.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	resp.Body.Close()
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
	resp, err := http.Post((SERVER_ROUTE + "/users"), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		test.Error(err)
	}

	if resp.StatusCode == 0 {
		req, err := http.NewRequest(http.MethodPost, (SERVER_ROUTE + "/login"), bytes.NewBuffer(reqBody))
		if err != nil {
			test.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			test.Fatal(err)
		}

		defer resp.Body.Close()

		// Check that the response status code is 200
		if resp.StatusCode != 200 {
			test.Errorf("handler returned wrong status code: got %v want %v",
				resp.StatusCode, http.StatusOK)
		}
		test.Log("Here is the cookie: " + string(Cookie))
		Cookie = resp.Header.Get("token")
		test.Log(string(Cookie))
		if Cookie != "" {
			test.Logf("PASSED. Cookie set")
		} else {
			test.Errorf("FAILED. Cookie not set")
		}

	}

}

func TestValidateToken(test *testing.T) {
	//Creating a new server for testing

	type Response struct {
		Name       string `json:"uname"`
		SessionExp string `json:"exp"`
	}
	var newSession *Response

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes.ValidateToken(w, r)
	}))
	defer ts.Close()

	reqString := map[string]string{
		"token": Cookie,
	}

	// Convert the map to a JSON-encoded byte array
	reqBody, err := json.Marshal(reqString)
	if err != nil {
		test.Errorf("Error converting map to JSON: %s", err)
	}

	fmt.Sprintln("request body: ", reqBody)

	resp, err := http.Post(SERVER_ROUTE+"/users", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		test.Error(err)
	}

	//If new user was made successfully. authenticate that user with valid credentials
	if resp.StatusCode == http.StatusOK { //WAS 0???
		req, err := http.NewRequest("ValidateToken", SERVER_ROUTE+"/token", bytes.NewBuffer(reqBody))
		if err != nil {
			test.Fatal(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			test.Fatal(err)
		}
		defer resp.Body.Close()

		// Check that the response status code is 200
		body, err := ioutil.ReadAll(resp.Body)
		respBody := json.Unmarshal(body, &newSession)

		if err != nil {
			test.Errorf("Error reading response body: %s", err)
		}

		if respBody == nil {
			test.Errorf("no session data found")
		}

		if newSession.SessionExp == "" {
			test.Errorf("no session data found")
		}

		if resp.StatusCode != http.StatusOK {
			test.Errorf("handler returned wrong status code: got %v want %v",
				resp.StatusCode, http.StatusOK)
		}
	}
}

//NOTE TestDELETE occurs after the tests on the session routes so that the test user can be used from those tests as well

// *********************************************routes.session TESTS****************************************//
func TestPostSession(t *testing.T) {
	// Create a new session object to send in the request body
	userID := TestUserID
	currTime := time.Now().Add(1 * time.Hour)

	// Create the JSON request string with placeholders for the variables
	jsonReqStr := `{"SessionID":1,"ArticleID":1,"UserID":%d,"CharHit":0,"CharMiss":0,"Time":"%s"}`

	// Fill in the placeholders with the actual values
	jsonReq := []byte(fmt.Sprintf(jsonReqStr, userID, currTime.Format(time.RFC3339)))

	// Create a new HTTP POST request with the JSON-encoded session in the body
	req, err := http.NewRequest("POST", SERVER_ROUTE+"/sessions", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the PostSession handler function and pass in the new request and response recorder
	handler := http.HandlerFunc(routes.PostSession)
	handler.ServeHTTP(rr, req)

	// Check the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body matches the original session object
	var responseSession models.Session
	err = json.Unmarshal(rr.Body.Bytes(), &responseSession)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetSession(t *testing.T) {
	req, err := http.NewRequest("GET", SERVER_ROUTE+"/session/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetSession)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusOK {
		t.Errorf("GetSession handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	/*
		expected := `{"SessionID":1,"Name":"Example Session","Description":"An example session"}`
		if rr.Body.String() != expected {
			t.Errorf("GetSession handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func TestGetSessionFromUser(t *testing.T) {
	req, err := http.NewRequest("GET", SERVER_ROUTE+"/session/user/"+fmt.Sprintf("%d", TestUserID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetSessionFromUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetSessionFromUser handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	/*
		expected := `{"SessionID":1,"Name":"Example Session","Description":"An example session"}`
		if rr.Body.String() != expected {
			t.Errorf("GetSessionFromUser handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func TestGetSessionFromArticle(t *testing.T) {
	req, err := http.NewRequest("GET", SERVER_ROUTE+"/articles/1/session", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetSessionFromArticle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetSessionFromArticle handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPutSession(t *testing.T) {
	// Create a new router and recorder
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()

	// Create a sample session
	session := models.Session{SessionID: 1, UserID: 1, ArticleID: 1}

	// Encode the session as JSON and create a request body
	requestBody, _ := json.Marshal(session)
	requestReader := bytes.NewReader(requestBody)

	// Create a PUT request with the session ID as a URL parameter
	request, _ := http.NewRequest("PUT", fmt.Sprintf("%s/sessions/%d", SERVER_ROUTE, session.SessionID), requestReader)

	// Set the content type header
	request.Header.Set("Content-Type", "application/json")

	// Attach the router to the recorder and serve the request
	router.HandleFunc("/sessions/{sid}", routes.PutSession)
	router.ServeHTTP(recorder, request)

	// Check the response code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	// Decode the response body and check if the updated session matches the original
	responseBody := recorder.Body.String()
	var updatedSession models.Session
	json.Unmarshal([]byte(responseBody), &updatedSession)
	if updatedSession.SessionID != session.SessionID || updatedSession.UserID != session.UserID || updatedSession.ArticleID != session.ArticleID {
		t.Errorf("Expected session %+v but got %+v", session, updatedSession)
	}
}

func TestDeleteSession(t *testing.T) {
	// Create a new router and recorder
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()

	// Create a sample session
	session := models.Session{SessionID: 1, UserID: 1, ArticleID: 1}

	// Save the session to the database
	initialize.DB.Create(&session)

	// Create a DELETE request with the session ID as a URL parameter
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/sessions/%d", SERVER_ROUTE, session.SessionID), nil)

	// Attach the router to the recorder and serve the request
	router.HandleFunc("/sessions/{sid}", routes.DeleteSession)
	router.ServeHTTP(recorder, request)

	// Check the response code
	if recorder.Code == http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	// Check if the session was deleted from the database
	var deletedSession models.Session
	initialize.DB.First(&deletedSession, session.SessionID)
	if deletedSession.SessionID != 0 {
		t.Errorf("Expected session to be deleted but found %+v", deletedSession)
	}
}

// *********************************************routes.article TESTS****************************************//

// TestDELETE is for routes.user.go
func TestDELETE(test *testing.T) {
	// Create a new test server with a handler function that handles the POST request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the method is POST
		if r.Method != "DELETE" {
			test.Errorf("Expected method DELETE, got %s", r.Method)
		}

		routes.DeleteUser(w, r)

	}))
	defer ts.Close()

	// Make a POST request to the test server
	reqBody := []byte(`{}`)
	req, err := http.NewRequest("DELETE", (SERVER_ROUTE + "/users/" + fmt.Sprintf("%d", TestUserID)), bytes.NewBuffer(reqBody))

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
func TestPostArticle(t *testing.T) {
	reqBody := []byte(`{"ID": "0", "Url": "https://en.wikipedia.org/wiki/Cat", "Length" : 200}`)
	req, err := http.NewRequest("POST", SERVER_ROUTE+"/article/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PostArticle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestGetArticle(t *testing.T) {
	req, err := http.NewRequest("GET", SERVER_ROUTE+"/article/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"aid": "1"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetArticle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestGetBody(t *testing.T) {
	req, err := http.NewRequest("GET", SERVER_ROUTE+"/article/body/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"aid": "1"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetBody)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestPutArticle(t *testing.T) {
	reqBody := []byte(`{"ID": "0", "Url": "https://en.wikipedia.org/wiki/Cat", "Length" : 400}`)
	req, err := http.NewRequest("PUT", SERVER_ROUTE+"/article/{aid}", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"aid": "1"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.PutArticle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDeleteArticle(t *testing.T) {
	req, err := http.NewRequest("DELETE", SERVER_ROUTE+"/article/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"sid": "1"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.DeleteArticle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
