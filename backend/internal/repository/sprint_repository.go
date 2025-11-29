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
	Create(title string) (*model.Sprint, error)
	Update(id int, title string, completed bool) (int, string, error)
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
	rows, err := r.db.Query("SELECT id, title, completed, created_at, updated_at FROM sprints WHERE is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sprints := []model.Sprint{}
	for rows.Next() {
		var s model.Sprint
		if err := rows.Scan(&s.ID, &s.Title, &s.Completed, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		sprints = append(sprints, s)
	}

	return sprints, nil
}

func (r *sprintRepository) Search(req *model.SprintSearchRequest) ([]model.Sprint, error) {
	query := "SELECT id, title, completed, created_at, updated_at FROM sprints WHERE is_deleted = false"
	args := []interface{}{}
	paramCount := 1

	// タイトルで部分一致検索
	if req.Title != nil {
		query += " AND title ILIKE $" + strconv.Itoa(paramCount)
		args = append(args, "%"+*req.Title+"%")
		paramCount++
	}

	// 完了状態でフィルタ
	if req.Completed != nil {
		query += " AND completed = $" + strconv.Itoa(paramCount)
		args = append(args, *req.Completed)
		paramCount++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sprints := []model.Sprint{}
	for rows.Next() {
		var s model.Sprint
		if err := rows.Scan(&s.ID, &s.Title, &s.Completed, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		sprints = append(sprints, s)
	}

	return sprints, nil
}

func (r *sprintRepository) Create(title string) (*model.Sprint, error) {
	s := &model.Sprint{
		Title:     title,
		Completed: false,
	}

	err := r.db.QueryRow(
		"INSERT INTO sprints (title) VALUES ($1) RETURNING id, created_at, updated_at",
		s.Title,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *sprintRepository) Update(id int, title string, completed bool) (int, string, error) {
	result, err := r.db.Exec(
		"UPDATE sprints SET title = $1, completed = $2, updated_at = NOW() WHERE id = $3 AND is_deleted = false",
		title, completed, id,
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

func (r *sprintRepository) Delete(id int) error {
	_, err := r.db.Exec(
		"UPDATE sprints SET is_deleted = true, updated_at = NOW() WHERE id = $1",
		id,
	)
	return err
}
