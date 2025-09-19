package domain

import "time"

// User represents the core user entity in the system
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Credential holds local authentication data for a user
type Credential struct {
	UserID       string `json:"user_id"`
	PasswordHash string `json:"-"` // Never serialize password hash
}

// SocialProfile links a user account to an external OAuth provider
type SocialProfile struct {
	UserID     string `json:"user_id"`
	Provider   string `json:"provider"`   // e.g., "google"
	ProviderID string `json:"provider_id"` // The unique ID from the provider
}