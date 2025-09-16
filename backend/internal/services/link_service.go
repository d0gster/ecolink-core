package services

import (
	"ecolink-core/internal/models"
	"ecolink-core/pkg/database"
	"ecolink-core/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/skip2/go-qrcode"
)

type LinkService struct {
	db      database.Database
	baseURL string
}

func NewLinkService(db database.Database, baseURL string) *LinkService {
	return &LinkService{
		db:      db,
		baseURL: baseURL,
	}
}

func (s *LinkService) CreateLink(originalURL, userID string) (*models.CreateLinkResponse, error) {
	// Check if link already exists for this URL and user
	userLinks, err := s.db.GetUserLinks(userID)
	if err == nil {
		for _, existingLink := range userLinks {
			if existingLink.URL == originalURL {
				// Return existent link
				shortURL := s.baseURL + "/" + existingLink.Code
				qrCode, err := s.generateQRCode(shortURL)
				if err != nil {
					return nil, err
				}
				return &models.CreateLinkResponse{
					ShortURL: shortURL,
					QRCode:   qrCode,
				}, nil
			}
		}
	}

	shortCode := utils.GenerateShortCode(originalURL)

	link := &models.Link{
		URL:       originalURL,
		Code:      shortCode,
		UserID:    userID,
		CreatedAt: time.Now(),
		Clicks:    0,
	}

	if err := s.db.SaveLink(link); err != nil {
		return nil, err
	}

	shortURL := s.baseURL + "/" + shortCode
	qrCode, err := s.generateQRCode(shortURL)
	if err != nil {
		return nil, err
	}

	return &models.CreateLinkResponse{
		ShortURL: shortURL,
		QRCode:   qrCode,
	}, nil
}

func (s *LinkService) GetOriginalURL(shortCode string) (string, error) {
	link, err := s.db.GetLink(shortCode)
	if err != nil {
		return "", err
	}

	// Increment click counter
	if err := s.db.IncrementClicks(shortCode); err != nil {
		return "", fmt.Errorf("failed to increment click counter: %w", err)
	}

	return link.URL, nil
}

func (s *LinkService) GetUserLinks(userID string) ([]*models.Link, error) {
	return s.db.GetUserLinks(userID)
}

func (s *LinkService) DeleteLink(code, userID string) error {
	// Check if link belongs to user
	link, err := s.db.GetLink(code)
	if err != nil {
		return err
	}

	if link.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.db.DeleteLink(code)
}

func (s *LinkService) generateQRCode(url string) (string, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(qr), nil
}
