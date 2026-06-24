package repositories

import (
	"gorm.io/gorm"
	"studyProject/entity"
)

type ChapterRepository interface {
	AddChapter(chapter entity.Chapter) error
	DeleteChapter(id int) error
	UpdateChapter(updateChapter entity.Chapter) error
	GetAllChapters() ([]entity.Chapter, error)
	GetChapter(id int) (entity.Chapter, error)
}

type chapterRepository struct {
	gormDB *gorm.DB
}

func NewChapterRepository(gormDB *gorm.DB) ChapterRepository {
	return &chapterRepository{gormDB: gormDB}
}

func (chapterRepository *chapterRepository) AddChapter(chapter entity.Chapter) error {
	result := chapterRepository.gormDB.Create(&chapter)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (chapterRepository *chapterRepository) DeleteChapter(id int) error {
	result := chapterRepository.gormDB.Delete(&entity.Chapter{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (chapterRepository *chapterRepository) UpdateChapter(updateChapter entity.Chapter) error {
	var chapter entity.Chapter
	check := chapterRepository.gormDB.First(&chapter, updateChapter.ID)

	if check.Error != nil {
		return check.Error
	}

	result := chapterRepository.gormDB.Save(&updateChapter)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (chapterRepository *chapterRepository) GetAllChapters() ([]entity.Chapter, error) {
	var chapters []entity.Chapter
	result := chapterRepository.gormDB.Find(&chapters)
	if result.Error != nil {
		return chapters, result.Error
	}
	return chapters, nil
}

func (chapterRepository *chapterRepository) GetChapter(id int) (entity.Chapter, error) {
	var chapter entity.Chapter
	result := chapterRepository.gormDB.First(&chapter, id)

	if result.Error != nil {
		return chapter, result.Error
	}
	return chapter, nil
}
