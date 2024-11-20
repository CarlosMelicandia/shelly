package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
  DiscordClientID    string
  DiscordClientSecretID    string
  DiscordRedirectURI    string
	JWTSecret          string
	JWTSecretRefresh   string
	TursoConnectionURL string
	TursoDatabaseName  string
	TursoAuthToken     string
}

func LoadEnv() AppConfig {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := AppConfig{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
    DiscordClientID:    os.Getenv("DISCORD_CLIENT_ID"),
    DiscordClientSecretID:    os.Getenv("DISCORD_CLIENT_SECRET"),
    DiscordRedirectURI:    os.Getenv("DISCORD_REDIRECT_URI"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTSecretRefresh:   os.Getenv("JWT_SECRET_REFRESH"),
		TursoConnectionURL: os.Getenv("TURSO_CONNECTION_URL"),
		TursoDatabaseName: os.Getenv("TURSO_DATABASE_NAME"),
    TursoAuthToken:     os.Getenv("TURSO_AUTH_TOKEN"),
	}

	if config.GoogleClientID == "" || 
     config.GoogleClientSecret == "" || 
     config.DiscordClientID == "" || 
     config.DiscordClientSecretID == "" || 
     config.DiscordRedirectURI == "" || 
     config.JWTSecret == "" || 
     config.JWTSecretRefresh == "" || 
     config.TursoConnectionURL == "" ||
     config.TursoDatabaseName == "" ||
     config.TursoAuthToken == "" {
      log.Fatal("Missing required environment variables")
	}

	return config
}
