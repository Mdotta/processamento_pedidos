package models

type ErrorResponse struct {
	Reason string `json:"reason"`
}

var (
	ErrEmailInUse     = ErrorResponse{Reason: "Email already in use"}
	ErrUserNotFound   = ErrorResponse{Reason: "User not found"}
	ErrInvalidRequest = ErrorResponse{Reason: "Invalid request"}
)
