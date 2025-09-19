package domain

import "time"

// AuthToken represents a JWT authentication token
type AuthToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    string    `json:"user_id"`
}

// TokenClaims represents the claims within a JWT token
type TokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
	Iat    int64  `json:"iat"`
}