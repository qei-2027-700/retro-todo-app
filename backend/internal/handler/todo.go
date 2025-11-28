package handler

import (
	"backend/internal/model"
	"backend/internal/repository"
	"net/http"
	"strconv"

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
	todos, err := h.repo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

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
	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createdTodo, err := h.repo.Create(t.Title, t.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdTodo)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	id, message, err := h.repo.Update(t.Title, t.Completed, t.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, strconv.Itoa(id)+message)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	idParam := c.Param("id") // パスパラメータ :id を取得
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter"})
	}

	err = h.repo.LogicalDelete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
