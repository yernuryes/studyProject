package services

import (
	"errors"
	"testing"

	"studyProject/entity"
	"studyProject/mocks"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddLesson(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	lesson := entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
		Content:     "Go language",
		Order:       1,
		ChapterID:   1,
	}

	mockRepo.
		On("AddLesson", lesson).
		Return(nil)

	err := service.AddLesson(lesson)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestAddLessonValidation(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	lesson := entity.Lesson{}

	err := service.AddLesson(lesson)

	assert.Error(t, err)

	assert.Equal(t, "invalid lesson data", err.Error())
}

func TestGetLesson(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	expectedLesson := entity.Lesson{
		ID:          1,
		Name:        "If-else Statement in Golang",
		Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
		Content:     "Go language",
		Order:       1,
		ChapterID:   1,
	}

	mockRepo.
		On("GetLesson", 1).
		Return(expectedLesson, nil)

	lesson, err := service.GetLesson(1)

	assert.NoError(t, err)

	assert.Equal(t, expectedLesson, lesson)

	mockRepo.AssertExpectations(t)
}

func TestGetLessonNotFound(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	mockRepo.
		On("GetLesson", 1).
		Return(entity.Lesson{}, gorm.ErrRecordNotFound)

	_, err := service.GetLesson(1)

	assert.Error(t, err)

	assert.Equal(t, gorm.ErrRecordNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteLesson(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	mockRepo.
		On("DeleteLesson", 1).
		Return(nil)

	err := service.DeleteLesson(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateLesson(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	lesson := entity.Lesson{
		ID:          1,
		Name:        "Updated",
		Description: "Updated desc",
		Content:     "Updated",
		Order:       1,
		ChapterID:   1,
	}

	mockRepo.
		On("UpdateLesson", lesson).
		Return(nil)

	err := service.UpdateLesson(lesson)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateLessonValidation(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	lesson := entity.Lesson{}

	err := service.UpdateLesson(lesson)

	assert.Error(t, err)

	assert.Equal(t, "invalid lesson data", err.Error())
}

func TestGetAllLessons(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	expectedLessons := []entity.Lesson{
		{
			ID:          1,
			Name:        "If-else Statement in Golang",
			Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
			Content:     "Go language",
			Order:       1,
			ChapterID:   1,
		},
	}

	mockRepo.
		On("GetAlllessons").
		Return(expectedLessons, nil)

	lessons, err := service.GetAllLessons()

	assert.NoError(t, err)

	assert.Equal(t, expectedLessons, lessons)

	mockRepo.AssertExpectations(t)
}

func TestRepositoryErrorThree(t *testing.T) {

	mockRepo := new(mocks.LessonRepository)

	service := NewLessonService(mockRepo)

	mockRepo.
		On("AddLesson", entity.Lesson{
			Name:        "If-else Statement in Golang",
			Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
			Content:     "Go language",
			Order:       1,
			ChapterID:   1,
		}).
		Return(errors.New("db error"))

	err := service.AddLesson(entity.Lesson{
		Name:        "If-else Statement in Golang",
		Description: "In Go (Golang), if-else statements provide basic conditional branching. Unlike many other languages, Go has specific syntactic rules—such as the mandatory use of curly braces and the omission of parentheses around conditions.",
		Content:     "Go language",
		Order:       1,
		ChapterID:   1,
	})

	assert.Error(t, err)

	assert.Equal(t, "db error", err.Error())

	mockRepo.AssertExpectations(t)
}
