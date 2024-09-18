package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CredentialsFile string
	TokenFile       string
}

func getEnv(key, fallback string) string {

	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return fallback

}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	credentialsPath := getEnv("CREDENTIALS_PATH", "config/credentials.json")
	tokenPath := getEnv("TOKEN_PATH", "config/token.json")

	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		log.Fatalf("Credentials file not found at path: %s", credentialsPath)
	}
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		log.Printf("Token file not found, it will be created at: %s", tokenPath)
	}

	return &Config{
		CredentialsFile: credentialsPath,
		TokenFile:       tokenPath,
	}
}
