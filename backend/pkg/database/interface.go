package database

import "ecolink-core/internal/models"

type Database interface {
	SaveLink(link *models.Link) error
	GetLink(code string) (*models.Link, error)
	GetUserLinks(userID string) ([]*models.Link, error)
	IncrementClicks(code string) error
	DeleteLink(code string) error
	SaveUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetUserByGoogleID(googleID string) (*models.User, error)
}