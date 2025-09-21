package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	NewUserID uuid.UUID `json:"newUserId"`
}

// if I need to hide a user value on response, I can change this struct
type GetUserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func (u User) ToGetUserResponse() GetUserResponse {
	return GetUserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
