package model

import "backend/internal/types"

type Sprint struct {
	ID        int              `json:"id"`
	Title     string           `json:"title"`
	Completed bool             `json:"completed"`
	CreatedAt types.CustomTime `json:"created_at"`
	UpdatedAt types.CustomTime `json:"updated_at"`
}

type SprintSearchRequest struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}
