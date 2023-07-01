package utilitis

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Getenv(key, fallback string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
