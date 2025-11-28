package repository

import (
	"backend/internal/model"
	"database/sql"
)

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
	Create(title string, completed bool) (*model.Todo, error)
	Update(title string, completed bool, id int) (int, string, error)
	LogicalDelete(id int) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, completed FROM todos WHERE is_deleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *todoRepository) Create(title string, completed bool) (*model.Todo, error) {
	t := &model.Todo{
		Title:     title,
		Completed: completed,
	}

	err := r.db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id",
		t.Title,
		t.Completed,
	).Scan(&t.ID)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) Update(title string, completed bool, id int) (int, string, error) {
	query := `UPDATE todos SET title = $1, completed = $2 WHERE id = $3`
	_, err := r.db.Exec(query, title, completed, id)
	if err != nil {
		return 0, "", err
	}
	return id, "更新成功", nil
}

func (r *todoRepository) LogicalDelete(id int) error {
	query := `UPDATE todos SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
