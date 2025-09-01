package models

import "time"

type Link struct {
	URL       string    `json:"original_url" firestore:"url"`
	Code      string    `json:"short_code" firestore:"code"`
	UserID    string    `json:"user_id" firestore:"user_id"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	Clicks    int       `json:"click_count" firestore:"clicks"`
}

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateLinkResponse struct {
	ShortURL string `json:"short_url"`
	QRCode   string `json:"qr_code"`
}