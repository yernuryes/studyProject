package repositories

import (
	"gorm.io/gorm"
	"studyProject/entity"
)

type LessonRepository interface {
	AddLesson(lesson entity.Lesson) error
	DeleteLesson(id int) error
	UpdateLesson(updateLesson entity.Lesson) error
	GetAllLessons() ([]entity.Lesson, error)
	GetLesson(id int) (entity.Lesson, error)
}

type lessonRepository struct {
	gormDB *gorm.DB
}

func NewLessonRepository(gormDB *gorm.DB) LessonRepository {
	return &lessonRepository{gormDB: gormDB}
}

func (lessonRepository *lessonRepository) AddLesson(lesson entity.Lesson) error {
	result := lessonRepository.gormDB.Create(&lesson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lessonRepository *lessonRepository) DeleteLesson(id int) error {
	result := lessonRepository.gormDB.Delete(&entity.Lesson{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (lessonRepository *lessonRepository) UpdateLesson(updateLesson entity.Lesson) error {
	var lesson entity.Lesson
	check := lessonRepository.gormDB.First(&lesson, updateLesson.ID)

	if check.Error != nil {
		return check.Error
	}

	result := lessonRepository.gormDB.Save(&updateLesson)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lessonRepository *lessonRepository) GetAllLessons() ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	result := lessonRepository.gormDB.Find(&lessons)
	if result.Error != nil {
		return lessons, result.Error
	}
	return lessons, nil
}

func (lessonRepository *lessonRepository) GetLesson(id int) (entity.Lesson, error) {
	var lesson entity.Lesson
	result := lessonRepository.gormDB.First(&lesson, id)

	if result.Error != nil {
		return lesson, result.Error
	}
	return lesson, nil
}
