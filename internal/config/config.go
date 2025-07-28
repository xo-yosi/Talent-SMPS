package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	DatabaseURL string
	S3Endpoint  string
	S3AccessKey string
	S3SecretKey string
	S3Region    string
}

var AppConfig Config

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using environment variables.")
	}
	env := os.Getenv("PRODENV")
	var dbHost, dbPort, dbUser, dbPassword, dbName string

	if env == "false" {
		dbHost = os.Getenv("POSTGRES_HOST")
		dbPort = os.Getenv("POSTGRES_PORT")
		dbUser = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbName = os.Getenv("POSTGRES_DB")
	} else {
		dbHost = os.Getenv("RENDER_HOST")
		dbPort = os.Getenv("RENDER_PORT")
		dbUser = os.Getenv("RENDER_USER")
		dbPassword = os.Getenv("RENDER_PASSWORD")
		dbName = os.Getenv("RENDER_DB")
	}

	appPort := os.Getenv("APP_PORT")
	s3Endpoint := os.Getenv("S3_ENDPOINT")
	s3AccessKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Region := os.Getenv("S3_REGION")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
		return Config{}, fmt.Errorf("database environment variables not set")
	}

	if appPort == "" {
		appPort = "8080"
	}

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	AppConfig := Config{
		AppPort:     appPort,
		DatabaseURL: databaseURL,
		S3Endpoint:  s3Endpoint,
		S3AccessKey: s3AccessKey,
		S3SecretKey: s3SecretKey,
		S3Region:    s3Region,
	}
	return AppConfig, nil
}
