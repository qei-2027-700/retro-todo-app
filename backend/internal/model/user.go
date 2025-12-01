package model

import "backend/internal/types"

type User struct {
	ID           int              `json:"id"`
	Username     string           `json:"username"`
	Email        string           `json:"email"`
	PasswordHash string           `json:"-"` // JSONには含めない
	ExternalID   *string          `json:"external_id,omitempty"`
	Provider     string           `json:"provider"`
	IsActive     bool             `json:"is_active"`
	CreatedAt    types.CustomTime `json:"created_at"`
	UpdatedAt    types.CustomTime `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
