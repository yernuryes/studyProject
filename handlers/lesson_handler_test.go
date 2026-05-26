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

func setupLessonRouter(handler LessonHandler) *gin.Engine {

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.Use(middleware.ErrorMiddleware())

	router.GET("/lessons", handler.HandleLessonGet)
	router.POST("/lessons", handler.HandleLessonPost)
	router.PUT("/lessons", handler.HandleLessonPut)
	router.DELETE("/lessons", handler.HandleLessonDelete)

	return router
}

func TestHandleLessonGet(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("GetAllLessons").
		Return([]entity.Lesson{}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/lessons",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleLessonGetByID(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("GetLesson", 1).
		Return(entity.Lesson{
			ID:          1,
			Name:        "If-else Statement in Golang",
			Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
			Content:     "Go language",
			Order:       1,
			ChapterID:   1,
		}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/lessons?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleLessonGetNotFound(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("GetLesson", 1).
		Return(entity.Lesson{}, gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/lessons?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleLessonPost(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	body := `{
		"name":"If-else Statement in Golang",
		"description":"In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
		"content":"Go language",
		"order":1,
		"chapter_id":1
	}`

	mockService.
		On("AddLesson", entity.Lesson{
			Name:        "If-else Statement in Golang",
			Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
			Content:     "Go language",
			Order:       1,
			ChapterID:   1,
		}).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/lessons",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandleLessonPostInvalidJSON(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	body := `invalid json`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/lessons",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHandleLessonDelete(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("DeleteLesson", 1).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/lessons?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleLessonDeleteNotFound(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("DeleteLesson", 1).
		Return(gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/lessons?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleLessonInternalError(t *testing.T) {

	mockService := new(mocks.LessonService)

	handler := NewLessonHandler(mockService)

	router := setupLessonRouter(handler)

	mockService.
		On("GetAllLessons").
		Return(nil, errors.New("db error"))

	req, _ := http.NewRequest(
		http.MethodGet,
		"/lessons",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
