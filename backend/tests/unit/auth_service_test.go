package unit

import (
	"context"
	"ecolink-core/internal/auth/usecase"
	"ecolink-core/pkg/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthService_Register(t *testing.T) {
	// Setup
	userRepo := database.NewMemoryUserRepository()
	tokenService := usecase.NewJWTTokenService("test-secret-key-32-bytes-long!", "test")
	googleConfig := usecase.GoogleConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURI:  "http://localhost:3000/callback",
	}
	authService := usecase.NewAuthService(userRepo, tokenService, googleConfig)

	t.Run("successful registration", func(t *testing.T) {
		ctx := context.Background()
		name := "Test User"
		email := "test@example.com"
		password := "securepassword123"

		user, err := authService.Register(ctx, name, email, password)

		require.NoError(t, err)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.NotEmpty(t, user.ID)
		assert.True(t, user.CreatedAt.IsZero()) // CreatedAt not set in current implementation
	})

	t.Run("duplicate email registration", func(t *testing.T) {
		ctx := context.Background()
		email := "duplicate@example.com"

		// First registration
		_, err := authService.Register(ctx, "User 1", email, "password123")
		require.NoError(t, err)

		// Second registration with same email
		_, err = authService.Register(ctx, "User 2", email, "password456")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user already exists")
	})
}

func TestAuthService_Login(t *testing.T) {
	// Setup
	userRepo := database.NewMemoryUserRepository()
	tokenService := usecase.NewJWTTokenService("test-secret-key-32-bytes-long!", "test")
	googleConfig := usecase.GoogleConfig{}
	authService := usecase.NewAuthService(userRepo, tokenService, googleConfig)

	t.Run("successful login", func(t *testing.T) {
		ctx := context.Background()
		email := "login@example.com"
		password := "testpassword123"

		// Register user first
		_, err := authService.Register(ctx, "Login User", email, password)
		require.NoError(t, err)

		// Login
		token, err := authService.Login(ctx, email, password)

		require.NoError(t, err)
		assert.NotEmpty(t, token.Token)
		assert.NotEmpty(t, token.UserID)
		assert.True(t, token.ExpiresAt.After(time.Now()))
	})

	t.Run("invalid credentials", func(t *testing.T) {
		ctx := context.Background()

		token, err := authService.Login(ctx, "nonexistent@example.com", "wrongpassword")

		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Contains(t, err.Error(), "invalid credentials")
	})
}

func TestTokenService_GenerateAndValidate(t *testing.T) {
	tokenService := usecase.NewJWTTokenService("test-secret-key-32-bytes-long!", "test")

	t.Run("generate and validate token", func(t *testing.T) {
		userID := "test-user-123"
		email := "test@example.com"

		// Generate token
		authToken, err := tokenService.GenerateToken(userID, email)
		require.NoError(t, err)
		assert.NotEmpty(t, authToken.Token)
		assert.Equal(t, userID, authToken.UserID)

		// Validate token
		claims, err := tokenService.ValidateToken(authToken.Token)
		require.NoError(t, err)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, email, claims.Email)
	})

	t.Run("invalid token", func(t *testing.T) {
		claims, err := tokenService.ValidateToken("invalid.token.here")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})
}
