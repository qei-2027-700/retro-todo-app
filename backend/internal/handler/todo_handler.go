package handler

//go:generate mockgen -source=todo_handler.go -destination=mock/mock_todo_handler.go -package=mock

import (
	"backend/internal/model"
	"backend/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandlerInterface interface {
	GetTodos(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
	SearchTodos(c echo.Context) error
}

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) TodoHandlerInterface {
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

	createdTodo, err := h.repo.Create(t.Title, t.Description, t.SprintID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdTodo)
}

// UpdateTodo godoc
// @Summary TODOを更新
// @Description 指定されたIDのTODOを更新します
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "TODO ID"
// @Param todo body model.Todo true "更新内容"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	// パスパラメータからIDを取得
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter"})
	}

	// ボディから更新内容を取得
	t := new(model.Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	rowsAffected, message, err := h.repo.Update(t.Title, t.Completed, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"rows_affected": rowsAffected,
		"message":       message,
	})
}

// SearchTodos godoc
// @Summary TODOを検索
// @Description 検索条件に基づいてTODOを検索します
// @Tags todos
// @Accept json
// @Produce json
// @Param request body model.TodoSearchRequest true "検索条件"
// @Success 200 {array} model.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/search [post]
func (h *TodoHandler) SearchTodos(c echo.Context) error {
	req := new(model.TodoSearchRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	todos, err := h.repo.Search(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todos)
}

// DeleteTodo godoc
// @Summary TODOを削除
// @Description 指定されたIDのTODOを論理削除します
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "TODO ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	idParam := c.Param("id") // パスパラメータ :id を取得
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID parameter"})
	}

	err = h.repo.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
