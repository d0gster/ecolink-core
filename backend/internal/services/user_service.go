package services

import (
	"ecolink-core/internal/models"
	"ecolink-core/pkg/database"
	"time"
	"crypto/rand"
	"encoding/hex"
)

type UserService struct {
	db database.Database
}

func NewUserService(db database.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateOrUpdateUser(googleID, name, email, picture string) (*models.User, error) {
	// Verifica se usuário já existe
	existingUser, err := s.db.GetUserByGoogleID(googleID)
	if err == nil {
		// Atualiza dados existentes
		existingUser.Name = name
		existingUser.Email = email
		existingUser.Picture = picture
		existingUser.UpdatedAt = time.Now()
		
		if err := s.db.SaveUser(existingUser); err != nil {
			return nil, err
		}
		return existingUser, nil
	}

	// Cria novo usuário
	userID := generateUserID()
	user := &models.User{
		ID:        userID,
		GoogleID:  googleID,
		Name:      name,
		Email:     email,
		Picture:   picture,
		Provider:  "google",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.db.GetUser(id)
}

func generateUserID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}