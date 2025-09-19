package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	BaseURL     string
	FrontendURL string
	Database    DatabaseConfig
	GoogleAuth  GoogleAuthConfig
	Security    SecurityConfig
	Cookie      CookieConfig
}

type DatabaseConfig struct {
	Type        string
	ProjectID   string
	Credentials string
}

type GoogleAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type SecurityConfig struct {
	JWTSecret    string
	CSRFSecret   string
	BcryptCost   int
	RateLimitRPS int
}

type CookieConfig struct {
	Domain   string
	Secure   bool
	SameSite string
}

func Load() (*Config, error) {
	godotenv.Load()

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
		Database: DatabaseConfig{
			Type:        getEnv("DB_TYPE", "memory"),
			ProjectID:   getEnv("FIRESTORE_PROJECT_ID", ""),
			Credentials: getEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
		},
		GoogleAuth: GoogleAuthConfig{
			ClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		},
		Security: SecurityConfig{
			JWTSecret:    os.Getenv("JWT_SECRET"),
			CSRFSecret:   getEnv("CSRF_SECRET", "change-me-in-production-32-chars-min"),
			BcryptCost:   getEnvInt("BCRYPT_COST", 12),
			RateLimitRPS: getEnvInt("RATE_LIMIT_RPS", 100),
		},
		Cookie: CookieConfig{
			Domain:   getEnv("COOKIE_DOMAIN", "localhost"),
			Secure:   getEnvAsBool("COOKIE_SECURE", false),
			SameSite: getEnv("COOKIE_SAMESITE", "Lax"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.Security.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	if len(c.Security.JWTSecret) < 32 {
		return fmt.Errorf("JWT_SECRET must be at least 32 characters")
	}

	if c.GoogleAuth.ClientID == "" || c.GoogleAuth.ClientSecret == "" {
		return fmt.Errorf("Google OAuth credentials are required")
	}

	validSameSite := []string{"strict", "lax", "none"}
	if !contains(validSameSite, strings.ToLower(c.Cookie.SameSite)) {
		return fmt.Errorf("invalid COOKIE_SAMESITE value: %s", c.Cookie.SameSite)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true"
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
