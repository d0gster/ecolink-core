package repository

import (
	"context"
	"ecolink-core/internal/auth/domain"
	"errors"
	"sync"
)

type InMemoryUserRepository struct {
	mu         sync.RWMutex
	users      map[string]*domain.User
	creds      map[string]*domain.Credential
	socials    map[string]*domain.SocialProfile
	nextUserID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:      make(map[string]*domain.User),
		creds:      make(map[string]*domain.Credential),
		socials:    make(map[string]*domain.SocialProfile),
		nextUserID: 1,
	}
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, *domain.Credential, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			cred := r.creds[user.ID]
			return user, cred, nil
		}
	}
	return nil, nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) FindByProviderID(ctx context.Context, provider, providerID string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, social := range r.socials {
		if social.Provider == provider && social.ProviderID == providerID {
			return r.users[social.UserID], nil
		}
	}
	return nil, errors.New("social profile not found")
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user *domain.User, cred *domain.Credential) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return errors.New("user already exists")
	}
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return errors.New("email already in use")
		}
	}

	r.users[user.ID] = user
	r.creds[cred.UserID] = cred
	return nil
}

func (r *InMemoryUserRepository) CreateUserFromSocial(ctx context.Context, user *domain.User, social *domain.SocialProfile) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; ok {
		return nil, errors.New("user already exists")
	}
	for _, existingSocial := range r.socials {
		if existingSocial.Provider == social.Provider && existingSocial.ProviderID == social.ProviderID {
			return nil, errors.New("social profile already in use")
		}
	}

	r.users[user.ID] = user
	r.socials[social.UserID] = social
	return user, nil
}

func (r *InMemoryUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}
