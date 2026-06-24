package repositories

import (
	"gorm.io/gorm"
	"studyProject/entity"
)

type CourseRepository interface {
	AddCourse(course entity.Course) error
	DeleteCourse(id int) error
	UpdateCourse(updateCourse entity.Course) error
	GetAllCourses() ([]entity.Course, error)
	GetCourse(id int) (entity.Course, error)
}

type courseRepository struct {
	gormDB *gorm.DB
}

func NewCourseRepository(gormDB *gorm.DB) CourseRepository {
	return &courseRepository{gormDB: gormDB}
}

func (courseRepository *courseRepository) AddCourse(course entity.Course) error {
	result := courseRepository.gormDB.Create(&course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (courseRepository *courseRepository) DeleteCourse(id int) error {
	result := courseRepository.gormDB.Delete(&entity.Course{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (courseRepository *courseRepository) UpdateCourse(updateCourse entity.Course) error {
	var course entity.Course

	check := courseRepository.gormDB.First(&course, updateCourse.ID)

	if check.Error != nil {
		return check.Error
	}

	result := courseRepository.gormDB.Save(&updateCourse)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (courseRepository *courseRepository) GetAllCourses() ([]entity.Course, error) {
	var courses []entity.Course
	result := courseRepository.gormDB.Find(&courses)
	if result.Error != nil {
		return courses, result.Error
	}
	return courses, nil
}

func (courseRepository *courseRepository) GetCourse(id int) (entity.Course, error) {
	var course entity.Course
	result := courseRepository.gormDB.First(&course, id)

	if result.Error != nil {
		return course, result.Error
	}
	return course, nil
}
