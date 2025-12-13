package repository

//go:generate mockgen -source=sprint_repository.go -destination=mock/mock_sprint_repository.go -package=mock

import (
	"backend/internal/model"
	"database/sql"
	"strconv"
)

type SprintRepository interface {
	FindAll() ([]model.Sprint, error)
	Search(req *model.SprintSearchRequest) ([]model.Sprint, error)
	Create(name, color string, isFavorite bool) (*model.Sprint, error)
	Update(id int, name, color string) (int, string, error)
	UpdateFavorite(id int, isFavorite bool) error
	Delete(id int) error
}

type sprintRepository struct {
	db *sql.DB
}

func NewSprintRepository(db *sql.DB) SprintRepository {
	return &sprintRepository{
		db: db,
	}
}

func (r *sprintRepository) FindAll() ([]model.Sprint, error) {
	rows, err := r.db.Query("SELECT id, name, color, is_favorite, created_at, updated_at FROM sprints WHERE is_deleted = false ORDER BY is_favorite DESC, created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sprints := []model.Sprint{}
	for rows.Next() {
		var s model.Sprint
		if err := rows.Scan(&s.ID, &s.Name, &s.Color, &s.IsFavorite, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		sprints = append(sprints, s)
	}

	return sprints, nil
}

func (r *sprintRepository) Search(req *model.SprintSearchRequest) ([]model.Sprint, error) {
	query := "SELECT id, name, color, is_favorite, created_at, updated_at FROM sprints WHERE is_deleted = false"
	args := []interface{}{}
	paramCount := 1

	// 名前で部分一致検索
	if req.Name != nil {
		query += " AND name ILIKE $" + strconv.Itoa(paramCount)
		args = append(args, "%"+*req.Name+"%")
		paramCount++
	}

	// お気に入りでフィルタ
	if req.IsFavorite != nil {
		query += " AND is_favorite = $" + strconv.Itoa(paramCount)
		args = append(args, *req.IsFavorite)
		paramCount++
	}

	query += " ORDER BY is_favorite DESC, created_at DESC"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sprints := []model.Sprint{}
	for rows.Next() {
		var s model.Sprint
		if err := rows.Scan(&s.ID, &s.Name, &s.Color, &s.IsFavorite, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		sprints = append(sprints, s)
	}

	return sprints, nil
}

func (r *sprintRepository) Create(name, color string, isFavorite bool) (*model.Sprint, error) {
	s := &model.Sprint{
		Name:       name,
		Color:      color,
		IsFavorite: isFavorite,
	}

	err := r.db.QueryRow(
		"INSERT INTO sprints (name, color, is_favorite) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		s.Name, s.Color, s.IsFavorite,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *sprintRepository) Update(id int, name, color string) (int, string, error) {
	result, err := r.db.Exec(
		"UPDATE sprints SET name = $1, color = $2, updated_at = NOW() WHERE id = $3 AND is_deleted = false",
		name, color, id,
	)
	if err != nil {
		return 0, "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, "", err
	}

	return int(rowsAffected), "Sprint updated successfully", nil
}

func (r *sprintRepository) UpdateFavorite(id int, isFavorite bool) error {
	_, err := r.db.Exec(
		"UPDATE sprints SET is_favorite = $1, updated_at = NOW() WHERE id = $2 AND is_deleted = false",
		isFavorite, id,
	)
	return err
}

func (r *sprintRepository) Delete(id int) error {
	_, err := r.db.Exec(
		"UPDATE sprints SET is_deleted = true, updated_at = NOW() WHERE id = $1",
		id,
	)
	return err
}
