package handler

import (
	"backend/internal/model"
	"backend/internal/storage"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	log.Println("[GET /todos] Fetching all todos...")
	rows, err := storage.DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		log.Printf("[GET /todos] Error querying database: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			log.Printf("[GET /todos] Error scanning row: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		todos = append(todos, t)
	}
	log.Printf("[GET /todos] Successfully fetched %d todos", len(todos))
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	log.Println("[POST /todos] Creating new todo...")
	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		log.Printf("[POST /todos] Error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	log.Printf("[POST /todos] Request data: title=%s, completed=%v", t.Title, t.Completed)

	err := storage.DB.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id", t.Title, t.Completed).Scan(&t.ID)
	if err != nil {
		log.Printf("[POST /todos] Error inserting todo: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	log.Printf("[POST /todos] Successfully created todo with ID: %d", t.ID)
	return c.JSON(http.StatusCreated, t)
}
