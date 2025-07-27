package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	DatabaseURL string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using environment variables.")
	}

	dbHost := os.Getenv("RENDER_HOST")
	dbPort := os.Getenv("RENDER_PORT")
	dbUser := os.Getenv("RENDER_USER")
	dbPassword := os.Getenv("RENDER_PASSWORD")
	dbName := os.Getenv("RENDER_DB")
	appPort := os.Getenv("APP_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
		return Config{}, fmt.Errorf("database environment variables not set")
	}

	if appPort == "" {
		appPort = "8080"
	}

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	return Config{
		AppPort:     appPort,
		DatabaseURL: databaseURL,
	}, nil
}
