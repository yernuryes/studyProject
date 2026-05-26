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

func setupCourseRouter(handler CourseHandler) *gin.Engine {

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.Use(middleware.ErrorMiddleware())

	router.GET("/courses", handler.HandleCourseGet)
	router.POST("/courses", handler.HandleCoursePost)
	router.PUT("/courses", handler.HandleCoursePut)
	router.DELETE("/courses", handler.HandleCourseDelete)

	return router
}

func TestHandleCourseGet(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("GetAllCourses").
		Return([]entity.Course{}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/courses",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleCourseGetByID(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("GetCourse", 1).
		Return(entity.Course{
			ID:          1,
			Name:        "Go",
			Description: "Backend",
		}, nil)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/courses?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleCourseGetNotFound(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("GetCourse", 1).
		Return(entity.Course{}, gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/courses?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleCoursePost(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	body := `{
		"name":"Go",
		"description":"Backend"
	}`

	mockService.
		On("AddCourse", entity.Course{
			Name:        "Go",
			Description: "Backend",
		}).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/courses",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestHandleCoursePostInvalidJSON(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	body := `invalid json`

	req, _ := http.NewRequest(
		http.MethodPost,
		"/courses",
		bytes.NewBuffer([]byte(body)),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHandleCourseDelete(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("DeleteCourse", 1).
		Return(nil)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/courses?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHandleCourseDeleteNotFound(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("DeleteCourse", 1).
		Return(gorm.ErrRecordNotFound)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/courses?id=1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestHandleCourseInternalError(t *testing.T) {

	mockService := new(mocks.CourseService)

	handler := NewCourseHandler(mockService)

	router := setupCourseRouter(handler)

	mockService.
		On("GetAllCourses").
		Return(nil, errors.New("db error"))

	req, _ := http.NewRequest(
		http.MethodGet,
		"/courses",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
