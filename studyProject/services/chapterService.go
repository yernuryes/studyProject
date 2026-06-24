package services

import (
	"errors"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/repositories"
)

type ChapterService interface {
	AddChapter(chapter entity.Chapter) error
	DeleteChapter(id int) error
	UpdateChapter(updateChapter entity.Chapter) error
	GetAllChapters() ([]entity.Chapter, error)
	GetChapter(id int) (entity.Chapter, error)
}

type chapterService struct {
	chapterRepository repositories.ChapterRepository
}

func NewChapterService(chapterRepository repositories.ChapterRepository) ChapterService {
	return &chapterService{chapterRepository: chapterRepository}
}

func (chapterService *chapterService) AddChapter(chapter entity.Chapter) error {
	logger.Log.Debug("Validating chapter data")
	if chapter.Name == "" || chapter.Description == "" {
		return errors.New("invalid course data")
	}
	logger.Log.Debug("Chapter validation successful")
	return chapterService.chapterRepository.AddChapter(chapter)
}

func (chapterService *chapterService) DeleteChapter(id int) error {
	return chapterService.chapterRepository.DeleteChapter(id)
}

func (chapterService *chapterService) UpdateChapter(updateChapter entity.Chapter) error {
	logger.Log.Debug("Updating chapter validation")
	if updateChapter.Name == "" || updateChapter.Description == "" {
		return errors.New("invalid course data")
	}
	return chapterService.chapterRepository.UpdateChapter(updateChapter)
}

func (chapterService *chapterService) GetAllChapters() ([]entity.Chapter, error) {
	return chapterService.chapterRepository.GetAllChapters()
}

func (chapterService *chapterService) GetChapter(id int) (entity.Chapter, error) {
	return chapterService.chapterRepository.GetChapter(id)
}
