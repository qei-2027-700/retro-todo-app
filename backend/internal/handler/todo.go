package handler

import (
	"backend/internal/model"
	"backend/internal/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

// GetTodos godoc
// @Summary TODOリストを取得
// @Description すべてのTODOを取得します
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Failure 500 {object} map[string]string
// @Router /todos [get]
func (h *TodoHandler) GetTodos(c echo.Context) error {
	log.Println("[HANDLER GET /todos] Fetching all todos...")

	todos, err := h.repo.FindAll()
	if err != nil {
		log.Printf("[HANDLER GET /todos] Error fetching todos: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Printf("[HANDLER GET /todos] Successfully fetched %d todos", len(todos))
	return c.JSON(http.StatusOK, todos)
}

// CreateTodo godoc
// @Summary TODOを作成
// @Description 新しいTODOを作成します
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.Todo true "TODO情報"
// @Success 201 {object} model.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	log.Println("[HANDLER POST /todos] Creating new todo...")

	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		log.Printf("[HANDLER POST /todos] Error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	log.Printf("[HANDLER POST /todos] Request data: title=%s, completed=%v", t.Title, t.Completed)

	createdTodo, err := h.repo.Create(t.Title, t.Completed)
	if err != nil {
		log.Printf("[HANDLER POST /todos] Error creating todo: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Printf("[HANDLER POST /todos] Successfully created todo with ID: %d", createdTodo.ID)
	return c.JSON(http.StatusCreated, createdTodo)
}
