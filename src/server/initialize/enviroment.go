package initialize

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error: Failed to load .env. Check that it is in the right directory")
	}
}
