package model

import "backend/internal/types"

type Todo struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Completed   bool             `json:"completed"`
	SprintID    *int             `json:"sprint_id"`
	CreatedAt   types.CustomTime `json:"created_at"`
	UpdatedAt   types.CustomTime `json:"updated_at"`
}

type TodoSearchRequest struct {
	Title       *string `json:"title"`       // 部分一致検索（任意）
	Description *string `json:"description"` // 部分一致検索（任意）
	Completed   *bool   `json:"completed"`   // 完了状態でフィルタ（任意）
	SprintID    *int    `json:"sprint_id"`   // スプリントIDでフィルタ（任意）
}
