package services

import (
	"errors"
	"testing"

	"studyProject/entity"
	"studyProject/mocks"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAddCourse(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	course := entity.Course{
		Name:        "Go",
		Description: "Backend",
	}

	mockRepo.
		On("AddCourse", course).
		Return(nil)

	err := service.AddCourse(course)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestAddCourseValidation(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	course := entity.Course{}

	err := service.AddCourse(course)

	assert.Error(t, err)

	assert.Equal(t, "invalid course data", err.Error())
}

func TestGetCourse(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	expectedCourse := entity.Course{
		ID:          1,
		Name:        "Go",
		Description: "Backend",
	}

	mockRepo.
		On("GetCourse", 1).
		Return(expectedCourse, nil)

	course, err := service.GetCourse(1)

	assert.NoError(t, err)

	assert.Equal(t, expectedCourse, course)

	mockRepo.AssertExpectations(t)
}

func TestGetCourseNotFound(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	mockRepo.
		On("GetCourse", 1).
		Return(entity.Course{}, gorm.ErrRecordNotFound)

	_, err := service.GetCourse(1)

	assert.Error(t, err)

	assert.Equal(t, gorm.ErrRecordNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCourse(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	mockRepo.
		On("DeleteCourse", 1).
		Return(nil)

	err := service.DeleteCourse(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCourse(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	course := entity.Course{
		ID:          1,
		Name:        "Updated",
		Description: "Updated desc",
	}

	mockRepo.
		On("UpdateCourse", course).
		Return(nil)

	err := service.UpdateCourse(course)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCourseValidation(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	course := entity.Course{}

	err := service.UpdateCourse(course)

	assert.Error(t, err)

	assert.Equal(t, "invalid course data", err.Error())
}

func TestGetAllCourses(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	expectedCourses := []entity.Course{
		{
			ID:          1,
			Name:        "Go",
			Description: "Backend",
		},
	}

	mockRepo.
		On("GetAllCourses").
		Return(expectedCourses, nil)

	courses, err := service.GetAllCourses()

	assert.NoError(t, err)

	assert.Equal(t, expectedCourses, courses)

	mockRepo.AssertExpectations(t)
}

func TestRepositoryError(t *testing.T) {

	mockRepo := new(mocks.CourseRepository)

	service := NewCourseService(mockRepo)

	mockRepo.
		On("AddCourse", entity.Course{
			Name:        "Go",
			Description: "Backend",
		}).
		Return(errors.New("db error"))

	err := service.AddCourse(entity.Course{
		Name:        "Go",
		Description: "Backend",
	})

	assert.Error(t, err)

	assert.Equal(t, "db error", err.Error())

	mockRepo.AssertExpectations(t)
}
