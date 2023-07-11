package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load environment: %v", err)
	}

	return os.Getenv(key)
}
