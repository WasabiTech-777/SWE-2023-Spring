package routes

import (
	"strings"
	"testing"

	"github.com/WasabiTech-777/SWE-2023-Spring/models"
	"golang.org/x/crypto/bcrypt"
)

//GOLANG documentation for testing: https://go.dev/doc/tutorial/add-a-test

func TestGenerateHashedPassword(t *testing.T) int {
	//create user with name/pw for testing
	const PASS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var funcUser, testUser models.User
	funcUser.Pass = PASS
	testUser.Pass = PASS
	GenerateHashedPassword(&funcUser)
	testPass, err := bcrypt.GenerateFromPassword([]byte(testUser.Pass), HASHCOST)

	if strings.Compare(funcUser.Pass, string(testPass)) != 0 || err != nil {
		return -1
	}
	return 0
}
