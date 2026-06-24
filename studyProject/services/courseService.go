package services

import (
	"errors"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/repositories"
)

type CourseService interface {
	AddCourse(course entity.Course) error
	DeleteCourse(id int) error
	UpdateCourse(updateCourse entity.Course) error
	GetAllCourses() ([]entity.Course, error)
	GetCourse(id int) (entity.Course, error)
}

type courseService struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(courseRepository repositories.CourseRepository) CourseService {
	return &courseService{courseRepository: courseRepository}
}

func (courseService *courseService) AddCourse(course entity.Course) error {
	logger.Log.Debug("Validating course data")
	if course.Name == "" || course.Description == "" {
		return errors.New("invalid course data")
	}
	logger.Log.Debug("Course validation successful")
	return courseService.courseRepository.AddCourse(course)

}

func (courseService *courseService) DeleteCourse(id int) error {
	return courseService.courseRepository.DeleteCourse(id)
}

func (courseService *courseService) UpdateCourse(updateCourse entity.Course) error {
	logger.Log.Debug("Updating course validation")
	if updateCourse.Name == "" || updateCourse.Description == "" {
		return errors.New("invalid course data")
	}
	return courseService.courseRepository.UpdateCourse(updateCourse)
}

func (courseService *courseService) GetAllCourses() ([]entity.Course, error) {
	return courseService.courseRepository.GetAllCourses()
}

func (courseService *courseService) GetCourse(id int) (entity.Course, error) {
	return courseService.courseRepository.GetCourse(id)
}
