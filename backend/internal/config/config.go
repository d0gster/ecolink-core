package config

import "os"

type Config struct {
	Port     string
	BaseURL  string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Type string // "memory" for demo, "firestore" for production
}

func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
		Database: DatabaseConfig{
			Type: getEnv("DB_TYPE", "memory"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}