package models

import "time"

type Link struct {
	URL       string    `json:"originalUrl" firestore:"originalUrl"`
	Code      string    `json:"shortCode" firestore:"shortCode"`
	UserID    string    `json:"userId" firestore:"userId"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	Clicks    int       `json:"clickCount" firestore:"clickCount"`
}

type CreateLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateLinkResponse struct {
	ShortURL string `json:"shortUrl"`
	QRCode   string `json:"qrCode"`
}
