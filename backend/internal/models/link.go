package models

import "time"

type Link struct {
	ID          string    `json:"id" firestore:"id"`
	OriginalURL string    `json:"original_url" firestore:"original_url"`
	ShortCode   string    `json:"short_code" firestore:"short_code"`
	UserID      string    `json:"user_id" firestore:"user_id"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	ClickCount  int       `json:"click_count" firestore:"click_count"`
}

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateLinkResponse struct {
	ShortURL string `json:"short_url"`
	QRCode   string `json:"qr_code"`
}