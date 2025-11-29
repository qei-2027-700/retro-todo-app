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

func TestGetSprints_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/sprints", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return([]model.Sprint{
		{
			ID:        1,
			Title:     "Sprint 1",
			Completed: false,
			CreatedAt: types.CustomTime(time.Now()),
			UpdatedAt: types.CustomTime(time.Now()),
		},
	}, nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.GetSprints(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var sprints []model.Sprint
	json.Unmarshal(rec.Body.Bytes(), &sprints)
	assert.Equal(t, 1, len(sprints))
	assert.Equal(t, "Sprint 1", sprints[0].Title)
}

func TestGetSprints_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/sprints", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return(nil, errors.New("database error"))

	handler := NewSprintHandler(mockRepo)
	err := handler.GetSprints(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateSprint_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	sprintJSON := `{"title":"New Sprint"}`
	req := httptest.NewRequest(http.MethodPost, "/sprints", strings.NewReader(sprintJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().Create("New Sprint").Return(&model.Sprint{
		ID:        1,
		Title:     "New Sprint",
		Completed: false,
		CreatedAt: types.CustomTime(time.Now()),
		UpdatedAt: types.CustomTime(time.Now()),
	}, nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.CreateSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var sprint model.Sprint
	json.Unmarshal(rec.Body.Bytes(), &sprint)
	assert.Equal(t, "New Sprint", sprint.Title)
}

func TestCreateSprint_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	invalidJSON := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/sprints", strings.NewReader(invalidJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockSprintRepository(ctrl)
	handler := NewSprintHandler(mockRepo)
	err := handler.CreateSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateSprint_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	sprintJSON := `{"title":"Updated Sprint","completed":true}`
	req := httptest.NewRequest(http.MethodPut, "/sprints/1", strings.NewReader(sprintJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().Update(1, "Updated Sprint", true).Return(1, "Sprint updated successfully", nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.UpdateSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateSprint_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	sprintJSON := `{"title":"Updated Sprint"}`
	req := httptest.NewRequest(http.MethodPut, "/sprints/invalid", strings.NewReader(sprintJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockRepo := mock.NewMockSprintRepository(ctrl)
	handler := NewSprintHandler(mockRepo)
	err := handler.UpdateSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateSprint_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	sprintJSON := `{"title":"Updated Sprint"}`
	req := httptest.NewRequest(http.MethodPut, "/sprints/999", strings.NewReader(sprintJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().Update(999, "Updated Sprint", false).Return(0, "", nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.UpdateSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteSprint_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/sprints/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().Delete(1).Return(nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.DeleteSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteSprint_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/sprints/invalid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	mockRepo := mock.NewMockSprintRepository(ctrl)
	handler := NewSprintHandler(mockRepo)
	err := handler.DeleteSprint(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSearchSprints_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	searchJSON := `{"title":"sprint"}`
	req := httptest.NewRequest(http.MethodPost, "/sprints/search", strings.NewReader(searchJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	title := "sprint"
	mockRepo := mock.NewMockSprintRepository(ctrl)
	mockRepo.EXPECT().Search(&model.SprintSearchRequest{
		Title: &title,
	}).Return([]model.Sprint{
		{
			ID:        1,
			Title:     "Sprint 1",
			Completed: false,
			CreatedAt: types.CustomTime(time.Now()),
			UpdatedAt: types.CustomTime(time.Now()),
		},
	}, nil)

	handler := NewSprintHandler(mockRepo)
	err := handler.SearchSprints(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var sprints []model.Sprint
	json.Unmarshal(rec.Body.Bytes(), &sprints)
	assert.Equal(t, 1, len(sprints))
}

func TestSearchSprints_InvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	invalidJSON := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/sprints/search", strings.NewReader(invalidJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := mock.NewMockSprintRepository(ctrl)
	handler := NewSprintHandler(mockRepo)
	err := handler.SearchSprints(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
