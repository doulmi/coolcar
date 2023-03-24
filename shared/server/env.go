package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv(configPath string) {
	err := godotenv.Load(configPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
