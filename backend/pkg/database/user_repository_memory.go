package database

import (
	"context"
	"ecolink-core/internal/auth/domain"
	"errors"
	"sync"
)

type MemoryUserRepository struct {
	users       map[string]*domain.User
	credentials map[string]*domain.Credential
	socials     map[string]*domain.SocialProfile
	mu          sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:       make(map[string]*domain.User),
		credentials: make(map[string]*domain.Credential),
		socials:     make(map[string]*domain.SocialProfile),
	}
}

func (r *MemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, *domain.Credential, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			cred := r.credentials[user.ID]
			return user, cred, nil
		}
	}
	return nil, nil, errors.New("user not found")
}

func (r *MemoryUserRepository) FindByProviderID(ctx context.Context, provider, providerID string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, social := range r.socials {
		if social.Provider == provider && social.ProviderID == providerID {
			user := r.users[social.UserID]
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MemoryUserRepository) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MemoryUserRepository) CreateUser(ctx context.Context, user *domain.User, cred *domain.Credential) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	if cred != nil {
		r.credentials[user.ID] = cred
	}
	return nil
}

func (r *MemoryUserRepository) CreateUserFromSocial(ctx context.Context, user *domain.User, social *domain.SocialProfile) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	r.socials[user.ID] = social
	return user, nil
}

func (r *MemoryUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}