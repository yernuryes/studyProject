package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/services"
)

type LessonHandler interface {
	HandleLessonGet(c *gin.Context)    // //w-writer, r-request
	HandleLessonPost(c *gin.Context)   // //w-writer, r-request
	HandleLessonPut(c *gin.Context)    // //w-writer, r-request
	HandleLessonDelete(c *gin.Context) // //w-writer, r-request
}

type lessonHandler struct {
	lessonService services.LessonService
}

func NewLessonHandler(lessonService services.LessonService) LessonHandler {
	return &lessonHandler{lessonService: lessonService}
}

// HandleLessonGet godoc
//
//	@Summary		Get lesson
//	@Description	Get lesson by id or all lessons
//	@Tags			lessons
//	@Produce		json
//	@Param			id	query		int	false	"Lesson ID"
//	@Success		200	{array}		entity.Lesson
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/lessons [get]
func (lessonHandler *lessonHandler) HandleLessonGet(c *gin.Context) {
	var idStr string = c.Query("id")

	// GET ALL
	if idStr == "" {

		logger.Log.Info("Getting all lessons")

		lessons, err := lessonHandler.lessonService.GetAllLessons()
		if err != nil {
			c.Error(err)
			return
		}

		logger.Log.Debugf("Lessons count: %d", len(lessons))

		c.JSON(http.StatusOK, lessons)
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

	logger.Log.Info("Getting lesson by id")

	logger.Log.Debugf("Lesson ID: %d", id)

	lesson, err := lessonHandler.lessonService.GetLesson(id)
	if err != nil {
		c.Error(err)
		return
	}

	logger.Log.Debugf("Lesson data: %+v", lesson)

	c.JSON(http.StatusOK, lesson)
}

// HandleLessonPost godoc
//
//	@Summary		Create lesson
//	@Description	Create new lesson
//	@Tags			lessons
//	@Accept			json
//	@Produce		json
//	@Param			lesson	body		entity.Lesson	true	"Lesson"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/lessons [post]
func (lessonHandler *lessonHandler) HandleLessonPost(c *gin.Context) {
	var lesson entity.Lesson

	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	logger.Log.Info("Creating new lesson")

	logger.Log.Debugf("Lesson data: %+v", lesson)

	err = lessonHandler.lessonService.AddLesson(lesson)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "lesson created",
	})
}

// HandleLessonPut godoc
//
//	@Summary		Update lesson
//	@Description	Update existing lesson
//	@Tags			lessons
//	@Accept			json
//	@Produce		json
//	@Param			lesson	body		entity.Lesson	true	"Updated Lesson"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/lessons [put]
func (lessonHandler *lessonHandler) HandleLessonPut(c *gin.Context) {
	var updateLesson entity.Lesson
	err := c.ShouldBindJSON(&updateLesson)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	logger.Log.Info("Updating lesson")

	logger.Log.Debugf("Updated lesson data: %+v", updateLesson)

	err = lessonHandler.lessonService.UpdateLesson(updateLesson)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "lesson updated",
	})
}

// HandleLessonDelete godoc
//
//	@Summary		Delete lesson
//	@Description	Delete lesson by id
//	@Tags			lessons
//	@Produce		json
//	@Param			id	query		int	true	"Lesson ID"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/lessons [delete]
func (lessonHandler *lessonHandler) HandleLessonDelete(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	logger.Log.Info("Deleting lesson")

	logger.Log.Debugf("Deleting lesson ID: %d", id)

	err = lessonHandler.lessonService.DeleteLesson(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "lesson deleted",
	})
}
