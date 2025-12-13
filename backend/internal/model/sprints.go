package model

import "backend/internal/types"

type Sprint struct {
	ID         int              `json:"id" gorm:"primaryKey"`
	Name       string           `json:"name" gorm:"not null"`
	Color      string           `json:"color"`
	IsFavorite bool             `json:"is_favorite" gorm:"default:false"`
	CreatedAt  types.CustomTime `json:"created_at"`
	UpdatedAt  types.CustomTime `json:"updated_at"`
}

type SprintSearchRequest struct {
	Name       *string `json:"name"`
	IsFavorite *bool   `json:"is_favorite"`
}

type UpdateFavoriteRequest struct {
	IsFavorite bool `json:"is_favorite"`
}
