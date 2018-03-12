package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAppKey() string {
	appKey := os.Getenv("APP_KEY")

	if appKey == "" {
		log.Fatal("APP_KEY null")
	}

	return appKey
}
