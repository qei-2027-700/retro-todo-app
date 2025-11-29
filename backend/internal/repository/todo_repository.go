package repository

//go:generate mockgen -source=todo_repository.go -destination=mock/mock_todo_repository.go -package=mock

import (
	"backend/internal/model"
	"database/sql"
	"strconv"
)

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
	Search(req *model.TodoSearchRequest) ([]model.Todo, error)
	Create(title string, description string, sprintID *int) (*model.Todo, error)
	Update(title string, completed bool, id int) (int, string, error)
	Delete(id int) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed, sprint_id, created_at, updated_at FROM todos WHERE is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.SprintID, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *todoRepository) Create(title string, description string, sprintID *int) (*model.Todo, error) {
	t := &model.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		SprintID:    sprintID,
	}

	err := r.db.QueryRow(
		"INSERT INTO todos (title, description, sprint_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		t.Title,
		t.Description,
		t.SprintID,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) Update(title string, completed bool, id int) (int, string, error) {
	result, err := r.db.Exec(
		"UPDATE todos SET title = $1, completed = $2, updated_at = NOW() WHERE id = $3 AND is_deleted = false",
		title, completed, id,
	)
	if err != nil {
		return 0, "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, "", err
	}

	return int(rowsAffected), "Todo updated successfully", nil
}

func (r *todoRepository) Search(req *model.TodoSearchRequest) ([]model.Todo, error) {
	query := "SELECT id, title, description, completed, sprint_id, created_at, updated_at FROM todos WHERE is_deleted = false"
	args := []interface{}{}
	paramCount := 1

	// タイトルで部分一致検索
	if req.Title != nil {
		query += " AND title ILIKE $" + strconv.Itoa(paramCount)
		args = append(args, "%"+*req.Title+"%")
		paramCount++
	}

	// 説明で部分一致検索
	if req.Description != nil {
		query += " AND description ILIKE $" + strconv.Itoa(paramCount)
		args = append(args, "%"+*req.Description+"%")
		paramCount++
	}

	// 完了状態でフィルタ
	if req.Completed != nil {
		query += " AND completed = $" + strconv.Itoa(paramCount)
		args = append(args, *req.Completed)
		paramCount++
	}

	// スプリントIDでフィルタ
	if req.SprintID != nil {
		query += " AND sprint_id = $" + strconv.Itoa(paramCount)
		args = append(args, *req.SprintID)
		paramCount++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.SprintID, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *todoRepository) Delete(id int) error {
	query := `UPDATE todos SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
