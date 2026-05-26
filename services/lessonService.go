package services

import (
	"errors"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/repositories"
)

type LessonService interface {
	AddLesson(lesson entity.Lesson) error
	DeleteLesson(id int) error
	UpdateLesson(updateLesson entity.Lesson) error
	GetAllLessons() ([]entity.Lesson, error)
	GetLesson(id int) (entity.Lesson, error)
}

type lessonService struct {
	lessonRepository repositories.LessonRepository
}

func NewLessonService(lessonRepository repositories.LessonRepository) LessonService {
	return &lessonService{lessonRepository: lessonRepository}
}

func (lessonService *lessonService) AddLesson(lesson entity.Lesson) error {
	logger.Log.Debug("Validating lesson data")
	if lesson.Name == "" || lesson.Description == "" {
		return errors.New("invalid course data")
	}
	logger.Log.Debug("Lesson validation successful")
	return lessonService.lessonRepository.AddLesson(lesson)
}

func (lessonService *lessonService) DeleteLesson(id int) error {
	return lessonService.lessonRepository.DeleteLesson(id)
}

func (lessonService *lessonService) UpdateLesson(updateLesson entity.Lesson) error {
	logger.Log.Debug("Updating lesson validation")
	if updateLesson.Name == "" || updateLesson.Description == "" {
		return errors.New("invalid course data")
	}
	return lessonService.lessonRepository.UpdateLesson(updateLesson)
}

func (lessonService *lessonService) GetAllLessons() ([]entity.Lesson, error) {
	return lessonService.lessonRepository.GetAllLessons()
}

func (lessonService *lessonService) GetLesson(id int) (entity.Lesson, error) {
	return lessonService.lessonRepository.GetLesson(id)
}
