package routes

import (
	"strings"
	"testing"

	"github.com/WasabiTech-777/SWE-2023-Spring/src/server/models"
)

const HASHCOST int = 16

func TestGenerateHashedPassword(test *testing.T) {
	//create user with name/pw for testing
	const PASS = "abcdefg0123456789"
	var testUser models.User
	testUser.Pass = PASS
	GenerateHashedPassword(&testUser)
	result := HelperTestGenerateHashedPassword(&testUser, PASS)
	if result != 0 {
		test.Errorf("GenerateHashedPassword failed")
	}

}

func HelperTestGenerateHashedPassword(testUser *models.User, password string) int {
	if strings.Compare(password, string(testUser.Pass)) != 0 {
		return 0
	}
	return -1
}
