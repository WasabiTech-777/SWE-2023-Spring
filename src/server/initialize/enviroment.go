package initialize

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error: Failed to load .env. Check that it is in the right directory")
	} else {
		fmt.Println("Loaded .env file!")
	}
}
