package repository

import (
	"context"
	"ecolink-core/internal/auth/domain"
)

// UserRepository defines the persistence port for user-related data
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, *domain.Credential, error)
	FindByProviderID(ctx context.Context, provider, providerID string) (*domain.User, error)
	FindByID(ctx context.Context, userID string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User, cred *domain.Credential) error
	CreateUserFromSocial(ctx context.Context, user *domain.User, social *domain.SocialProfile) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}