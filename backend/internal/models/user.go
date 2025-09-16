package models

import "time"

type User struct {
	ID        string    `json:"id" firestore:"id"`
	GoogleID  string    `json:"googleId" firestore:"googleId"`
	Name      string    `json:"name" firestore:"name"`
	Email     string    `json:"email" firestore:"email"`
	Picture   string    `json:"picture" firestore:"picture"`
	Provider  string    `json:"provider" firestore:"provider"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" firestore:"updatedAt"`
}

type CreateUserRequest struct {
	GoogleID string `json:"googleId" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Picture  string `json:"picture"`
}
