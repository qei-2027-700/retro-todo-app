package repository

import (
	"backend/internal/model"
	"database/sql"
	"log"
)

type TodoRepository interface {
	FindAll() ([]model.Todo, error)
	Create(title string, completed bool) (*model.Todo, error)
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
	log.Println("[REPOSITORY] Fetching all todos from database...")

	rows, err := r.db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		log.Printf("[REPOSITORY] Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			log.Printf("[REPOSITORY] Error scanning row: %v", err)
			return nil, err
		}
		todos = append(todos, t)
	}

	log.Printf("[REPOSITORY] Successfully fetched %d todos", len(todos))
	return todos, nil
}

func (r *todoRepository) Create(title string, completed bool) (*model.Todo, error) {
	log.Printf("[REPOSITORY] Creating todo: title=%s, completed=%v", title, completed)

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
		log.Printf("[REPOSITORY] Error inserting todo: %v", err)
		return nil, err
	}

	log.Printf("[REPOSITORY] Successfully created todo with ID: %d", t.ID)
	return t, nil
}
