package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig this function loads the .env file.
func LoadConfig() {
	// Load .env file into environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file.")
	}
}

// GetEnviromentVariable this function loads an especific environment variable.
func GetEnviromentVariable(key string) string {
	value, exists := os.LookupEnv(key)
	// Checks if the environment variable exists.
	if !exists {
		log.Fatal("error finding the " + key + " environment variable.")
		return ""
	}
	return value
}
