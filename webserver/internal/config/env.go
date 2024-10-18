package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	JWTSecret          string
}

func LoadEnv() AppConfig {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := AppConfig{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
	}

	if config.GoogleClientID == "" || config.GoogleClientSecret == "" || config.JWTSecret == "" {
		log.Fatal("Missing required environment variables")
	}

	return config
}
