package services

import (
	"ecolink-core/internal/models"
	"ecolink-core/pkg/database"
	"ecolink-core/pkg/utils"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type LinkService struct {
	db      *database.MemoryDB
	baseURL string
}

func NewLinkService(db *database.MemoryDB, baseURL string) *LinkService {
	return &LinkService{
		db:      db,
		baseURL: baseURL,
	}
}

func (s *LinkService) CreateLink(originalURL, userID string) (*models.CreateLinkResponse, error) {
	shortCode := utils.GenerateShortCode(originalURL)
	
	link := &models.Link{
		ID:          uuid.New().String(),
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		UserID:      userID,
		CreatedAt:   time.Now(),
		ClickCount:  0,
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
	link, err := s.db.GetLinkByCode(shortCode)
	if err != nil {
		return "", err
	}
	
	// Incrementa contador de cliques
	s.db.IncrementClick(shortCode)
	
	return link.OriginalURL, nil
}

func (s *LinkService) GetUserLinks(userID string) ([]*models.Link, error) {
	return s.db.GetLinksByUser(userID)
}

func (s *LinkService) generateQRCode(url string) (string, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(qr), nil
}