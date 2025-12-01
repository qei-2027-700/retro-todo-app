package handler

//go:generate mockgen -source=sprint_handler.go -destination=mock/mock_sprint_handler.go -package=mock

import (
	"backend/internal/model"
	"backend/internal/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SprintHandlerInterface interface {
	GetSprints(c echo.Context) error
	CreateSprint(c echo.Context) error
	UpdateSprint(c echo.Context) error
	DeleteSprint(c echo.Context) error
	SearchSprints(c echo.Context) error
}

type SprintHandler struct {
	repo repository.SprintRepository
}

func NewSprintHandler(repo repository.SprintRepository) SprintHandlerInterface {
	return &SprintHandler{repo: repo}
}

// GetSprints godoc
// @Summary スプリントリストを取得
// @Description すべてのスプリントを取得します
// @Tags sprints
// @Accept json
// @Produce json
// @Success 200 {array} model.Sprint
// @Failure 500 {object} map[string]string
// @Router /sprints [get]
func (h *SprintHandler) GetSprints(c echo.Context) error {
	sprints, err := h.repo.FindAll()
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, sprints)
}

// CreateSprint godoc
// @Summary スプリントを作成
// @Description 新しいスプリントを作成します
// @Tags sprints
// @Accept json
// @Produce json
// @Param sprint body model.Sprint true "スプリント情報"
// @Success 201 {object} model.Sprint
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sprints [post]
func (h *SprintHandler) CreateSprint(c echo.Context) error {
	s := new(model.Sprint)
	if err := c.Bind(s); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	createdSprint, err := h.repo.Create(s.Title)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, createdSprint)
}

// SearchSprints godoc
// @Summary スプリントを検索
// @Description 検索条件に基づいてスプリントを検索します
// @Tags sprints
// @Accept json
// @Produce json
// @Param request body model.SprintSearchRequest true "検索条件"
// @Success 200 {array} model.Sprint
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sprints/search [post]
func (h *SprintHandler) SearchSprints(c echo.Context) error {
	req := new(model.SprintSearchRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	sprints, err := h.repo.Search(req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, sprints)
}

// UpdateSprint godoc
// @Summary スプリントを更新
// @Description 指定されたIDのスプリントを更新します
// @Tags sprints
// @Accept json
// @Produce json
// @Param id path int true "スプリント ID"
// @Param sprint body model.Sprint true "更新内容"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sprints/{id} [put]
func (h *SprintHandler) UpdateSprint(c echo.Context) error {
	// パスパラメータからIDを取得
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID parameter"})
	}

	// ボディから更新内容を取得
	s := new(model.Sprint)
	if err := c.Bind(s); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	rowsAffected, message, err := h.repo.Update(id, s.Title, s.Completed)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(404, map[string]string{"error": "Sprint not found"})
	}

	return c.JSON(200, map[string]interface{}{
		"rows_affected": rowsAffected,
		"message":       message,
	})
}

// DeleteSprint godoc
// @Summary スプリントを削除
// @Description 指定されたIDのスプリントを論理削除します
// @Tags sprints
// @Accept json
// @Produce json
// @Param id path int true "スプリント ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /sprints/{id} [delete]
func (h *SprintHandler) DeleteSprint(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid sprint ID"})
	}

	err = h.repo.Delete(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Sprint deleted successfully"})
}
