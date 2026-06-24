package services

import (
	"errors"
	"testing"

	"studyProject/entity"
	"studyProject/mocks"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddChapter(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	chapter := entity.Chapter{
		Name:        "Control structures",
		Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic. Unlike many other languages, Go lacks a dedicated while keyword, instead utilizing a versatile for loop for all iteration needs",
		Order:       1,
		CourseID:    1,
	}

	mockRepo.
		On("AddChapter", chapter).
		Return(nil)

	err := service.AddChapter(chapter)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestAddChapterValidation(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	chapter := entity.Chapter{}

	err := service.AddChapter(chapter)

	assert.Error(t, err)

	assert.Equal(t, "invalid chapter data", err.Error())
}

func TestGetChapter(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	expectedChapter := entity.Chapter{
		ID:          1,
		Name:        "Control structures",
		Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
		Order:       1,
		CourseID:    1,
	}

	mockRepo.
		On("GetChapter", 1).
		Return(expectedChapter, nil)

	chapter, err := service.GetChapter(1)

	assert.NoError(t, err)

	assert.Equal(t, expectedChapter, chapter)

	mockRepo.AssertExpectations(t)
}

func TestGetChapterNotFound(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	mockRepo.
		On("GetChapter", 1).
		Return(entity.Chapter{}, gorm.ErrRecordNotFound)

	_, err := service.GetChapter(1)

	assert.Error(t, err)

	assert.Equal(t, gorm.ErrRecordNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteChapter(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	mockRepo.
		On("DeleteChapter", 1).
		Return(nil)

	err := service.DeleteChapter(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateChapter(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	chapter := entity.Chapter{
		ID:          1,
		Name:        "Updated",
		Description: "Updated desc",
		Order:       1,
		CourseID:    1,
	}

	mockRepo.
		On("UpdateChapter", chapter).
		Return(nil)

	err := service.UpdateChapter(chapter)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateChapterValidation(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	chapter := entity.Chapter{}

	err := service.UpdateChapter(chapter)

	assert.Error(t, err)

	assert.Equal(t, "invalid chapter data", err.Error())
}

func TestGetAllChapters(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	expectedChapters := []entity.Chapter{
		{
			ID:          1,
			Name:        "Control structures",
			Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
			Order:       1,
			CourseID:    1,
		},
		{
			ID:          2,
			Name:        "Variables",
			Description: "In Go (Golang), variables are explicitly declared containers used by the compiler to store data and check for type correctness.",
			Order:       2,
			CourseID:    1,
		},
	}

	mockRepo.
		On("GetAllChapters").
		Return(expectedChapters, nil)

	chapters, err := service.GetAllChapters()

	assert.NoError(t, err)

	assert.Equal(t, expectedChapters, chapters)

	mockRepo.AssertExpectations(t)
}

func TestRepositoryErrorTwo(t *testing.T) {

	mockRepo := new(mocks.ChapterRepository)

	service := NewChapterService(mockRepo)

	mockRepo.
		On("AddChapter", entity.Chapter{
			Name:        "Control structures",
			Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
			Order:       1,
			CourseID:    1,
		}).
		Return(errors.New("db error"))

	err := service.AddChapter(entity.Chapter{
		Name:        "Control structures",
		Description: "Go uses a streamlined set of control structures—primarily if, for, and switch—to manage program logic.",
		Order:       1,
		CourseID:    1,
	})

	assert.Error(t, err)

	assert.Equal(t, "db error", err.Error())

	mockRepo.AssertExpectations(t)
}
