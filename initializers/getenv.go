package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func GetenvironmentVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
