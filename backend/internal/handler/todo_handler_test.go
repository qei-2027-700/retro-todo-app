package handler

import (
	"backend/internal/model"
	"backend/internal/repository/mock"
	"backend/internal/types"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTodos_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return([]model.Todo{
		{
			ID:          1,
			Title:       "Test Todo",
			Description: "Test Description",
			Completed:   false,
			SprintID:    nil,
			CreatedAt:   types.CustomTime(time.Now()),
			UpdatedAt:   types.CustomTime(time.Now()),
		},
	}, nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.GetTodos(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var todos []model.Todo
	json.Unmarshal(rec.Body.Bytes(), &todos)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, "Test Todo", todos[0].Title)
}

func TestGetTodos_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return(nil, errors.New("database error"))

	handler := NewTodoHandler(mockRepo)
	err := handler.GetTodos(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateTodo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	todoJSON := `{"title":"New Todo","description":"New Description"}`
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Create("New Todo", "New Description", (*int)(nil)).Return(&model.Todo{
		ID:          1,
		Title:       "New Todo",
		Description: "New Description",
		Completed:   false,
		SprintID:    nil,
		CreatedAt:   types.CustomTime(time.Now()),
		UpdatedAt:   types.CustomTime(time.Now()),
	}, nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.CreateTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var todo model.Todo
	json.Unmarshal(rec.Body.Bytes(), &todo)
	assert.Equal(t, "New Todo", todo.Title)
	assert.Equal(t, "New Description", todo.Description)
}

func TestCreateTodo_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	invalidJSON := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(invalidJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockTodoRepository(ctrl)
	handler := NewTodoHandler(mockRepo)
	err := handler.CreateTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateTodo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	todoJSON := `{"title":"Updated Todo","completed":true}`
	req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Update("Updated Todo", true, 1).Return(1, "Todo updated successfully", nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.UpdateTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateTodo_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	todoJSON := `{"title":"Updated Todo"}`
	req := httptest.NewRequest(http.MethodPut, "/todos/invalid", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockRepo := mock.NewMockTodoRepository(ctrl)
	handler := NewTodoHandler(mockRepo)
	err := handler.UpdateTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateTodo_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	todoJSON := `{"title":"Updated Todo"}`
	req := httptest.NewRequest(http.MethodPut, "/todos/999", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Update("Updated Todo", false, 999).Return(0, "", nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.UpdateTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteTodo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Delete(1).Return(nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.DeleteTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteTodo_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockRepo := mock.NewMockTodoRepository(ctrl)
	handler := NewTodoHandler(mockRepo)
	err := handler.DeleteTodo(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSearchTodos_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	searchJSON := `{"title":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/todos/search", strings.NewReader(searchJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	title := "test"
	mockRepo := mock.NewMockTodoRepository(ctrl)
	mockRepo.EXPECT().Search(&model.TodoSearchRequest{
		Title: &title,
	}).Return([]model.Todo{
		{
			ID:          1,
			Title:       "test todo",
			Description: "test description",
			Completed:   false,
			CreatedAt:   types.CustomTime(time.Now()),
			UpdatedAt:   types.CustomTime(time.Now()),
		},
	}, nil)

	handler := NewTodoHandler(mockRepo)
	err := handler.SearchTodos(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var todos []model.Todo
	json.Unmarshal(rec.Body.Bytes(), &todos)
	assert.Equal(t, 1, len(todos))
}
