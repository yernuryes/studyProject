package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/services"
)

type ChapterHandler interface {
	HandleChapterGet(c *gin.Context)    // //w-writer, r-request
	HandleChapterPost(c *gin.Context)   // //w-writer, r-request
	HandleChapterPut(c *gin.Context)    // //w-writer, r-request
	HandleChapterDelete(c *gin.Context) // //w-writer, r-request
}

type chapterHandler struct {
	chapterService services.ChapterService
}

func NewChapterHandler(chapterService services.ChapterService) ChapterHandler {
	return &chapterHandler{chapterService: chapterService}
}

// HandleChapterGet godoc
//
//	@Summary		Get chapter
//	@Description	Get chapter by id or all chapters
//	@Tags			chapters
//	@Produce		json
//	@Param			id	query		int	false	"Chapter ID"
//	@Success		200	{array}		entity.Chapter
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/chapters [get]
func (chapterHandler *chapterHandler) HandleChapterGet(c *gin.Context) {
	var idStr string = c.Query("id")

	// GET ALL
	if idStr == "" {
		logger.Log.Info("Getting all chapters")
		chapters, err := chapterHandler.chapterService.GetAllChapters()
		if err != nil {
			c.Error(err)
			return
		}

		logger.Log.Debugf("Chapters count: %d", len(chapters))

		c.JSON(http.StatusOK, chapters)
		return
	}

	// GET BY ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	logger.Log.Info("Getting chapter by id")

	logger.Log.Debugf("Course ID: %d", id)

	chapter, err := chapterHandler.chapterService.GetChapter(id)
	if err != nil {
		c.Error(err)
		return
	}

	logger.Log.Debugf("Chapter data: %+v", chapter)

	c.JSON(http.StatusOK, chapter)
}

// HandleChapterPost godoc
//
//	@Summary		Create chapter
//	@Description	Create new chapter
//	@Tags			chapters
//	@Accept			json
//	@Produce		json
//	@Param			chapter	body		entity.Chapter	true	"Chapter"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/chapters [post]
func (chapterHandler *chapterHandler) HandleChapterPost(c *gin.Context) {
	var chapter entity.Chapter

	err := c.ShouldBindJSON(&chapter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	logger.Log.Info("Creating new chapter")

	logger.Log.Debugf("Chapter data: %+v", chapter)

	err = chapterHandler.chapterService.AddChapter(chapter)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "chapter created",
	})
}

// HandleChapterPut godoc
//
//	@Summary		Update chapter
//	@Description	Update existing chapter
//	@Tags			chapters
//	@Accept			json
//	@Produce		json
//	@Param			chapter	body		entity.Chapter	true	"Updated Chapter"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/chapters [put]
func (chapterHandler *chapterHandler) HandleChapterPut(c *gin.Context) {
	var updateChapter entity.Chapter
	err := c.ShouldBindJSON(&updateChapter)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	logger.Log.Info("Updating chapter")

	logger.Log.Debugf("Updated chapter data: %+v", updateChapter)

	err = chapterHandler.chapterService.UpdateChapter(updateChapter)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "chapter updated",
	})
}

// HandleChapterDelete godoc
//
//	@Summary		Delete chapter
//	@Description	Delete chapter by id
//	@Tags			chapters
//	@Produce		json
//	@Param			id	query		int	true	"Chapter ID"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/chapters [delete]
func (chapterHandler *chapterHandler) HandleChapterDelete(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	logger.Log.Info("Deleting chapter")

	logger.Log.Debugf("Deleting chapter ID: %d", id)

	err = chapterHandler.chapterService.DeleteChapter(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chapter deleted",
	})
}
