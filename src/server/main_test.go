package main

import (
	"testing"
	"os"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"github.com/WasabiTech-777/SWE-2023-Spring/initialize"
	"github.com/WasabiTech-777/SWE-2023-Spring/routes"
)

func TestLoadEnv(test *testing.T) {
	initialize.LoadEnv()
	dsn := os.Getenv("DSNonline")
	port := os.Getenv("PORT")
	if dsn == "" || port == "" {
		test.Errorf("FAILED. DSN and PORT not sourced")
	} else{
		test.Logf("PASSED. DSN and PORT sourced")
	}	
}

func TestConnect (test *testing.T) {
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

func TestMigrate (test *testing.T) {
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