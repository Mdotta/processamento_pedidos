package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string // not exported in responses
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // not stored, just for creation
}

type CreateUserResponse struct {
	NewUserID uuid.UUID `json:"newUserId"`
}

// if I need to hide a user value on response, I can change this struct
type GetUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (u User) ToGetUserResponse() GetUserResponse {
	return GetUserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
