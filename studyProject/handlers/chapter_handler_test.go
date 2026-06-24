package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"studyProject/middleware"
	"testing"

	"studyProject/entity"
	"studyProject/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupChapterRouter(handler ChapterHandler) *gin.Engine {

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.Use(middleware.ErrorMiddleware())

	router.GET("/chapters", handler.HandleChapterGet)
	router.POST("/chapters", handler.HandleChapterPost)
	router.PUT("/chapters", handler.HandleChapterPut)
	router.DELETE("/chapters", handler.HandleChapterDelete)

	return router
}

func TestHandleChapterGet(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("GetAllChapters").
		Return([]entity.Chapter{}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/chapters",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleChapterGetByID(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("GetChapter", 1).
		Return(entity.Chapter{
			ID:          1,
			Name:        "Control structures",
			Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
			Order:       1,
			CourseID:    1,
		}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/chapters?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleChapterGetNotFound(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("GetChapter", 1).
		Return(entity.Chapter{}, gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/chapters?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleChapterPost(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	body := `{
		"name":"Control structures",
		"description":"Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
		"order":1,
		"courseID":1
	}`

	mockService.
		On("AddChapter", entity.Chapter{
			Name:        "Control structures",
			Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
			Order:       1,
			CourseID:    1,
		}).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/chapters",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandleChapterPostInvalidJSON(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	body := `invalid json`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/chapters",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHandleChapterDelete(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("DeleteChapter", 1).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/chapters?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleChapterDeleteNotFound(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("DeleteChapter", 1).
		Return(gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/chapters?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleChapterInternalError(t *testing.T) {

	mockService := new(mocks.ChapterService)

	handler := NewChapterHandler(mockService)

	router := setupChapterRouter(handler)

	mockService.
		On("GetAllChapters").
		Return(nil, errors.New("db error"))

	req, _ := http.NewRequest(
		http.MethodGet,
		"/chapters",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
