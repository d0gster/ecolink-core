package models

import "time"

type User struct {
	ID        string    `json:"id" firestore:"id"`
	GoogleID  string    `json:"google_id" firestore:"google_id"`
	Name      string    `json:"name" firestore:"name"`
	Email     string    `json:"email" firestore:"email"`
	Picture   string    `json:"picture" firestore:"picture"`
	Provider  string    `json:"provider" firestore:"provider"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at"`
}

type CreateUserRequest struct {
	GoogleID string `json:"google_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Picture  string `json:"picture"`
}