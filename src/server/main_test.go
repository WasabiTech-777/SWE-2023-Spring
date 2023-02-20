package main

import (
	"testing"
	"github.com/WasabiTech-777/SWE-2023-Spring/initialize"
)

func TestMain (m *testing.M) {
	initialize.LoadEnv()
	initialize.Connect()
	initialize.Migrate()
}

func tableExists(name string) string{
	sqlString := `SELECT EXISTS (
		SELECT FROM
			pg_tables
		WHERE
			schemaname = 'public' AND
			tablename  = ''
		);`
	return sqlString
}


func clearUsers () {
	initialize.DB.Exec("DELETE * FROM USERS")
}