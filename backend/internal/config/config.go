package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	BaseURL           string
	Database          DatabaseConfig
	Auth0Domain       string
	Auth0Audience     string
}

type DatabaseConfig struct {
	Type        string
	ProjectID   string
	Credentials string
}

func Load() *Config {
	godotenv.Load()
	
	return &Config{
		Port:    getEnv("PORT", "8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
		Database: DatabaseConfig{
			Type:        getEnv("DB_TYPE", "memory"),
			ProjectID:   getEnv("FIRESTORE_PROJECT_ID", ""),
			Credentials: getEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
		},
		Auth0Domain:   getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience: getEnv("AUTH0_AUDIENCE", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}