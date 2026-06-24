package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyProject/entity"
	"studyProject/logger"
	"studyProject/services"
)

type CourseHandler interface {
	HandleCourseGet(c *gin.Context)    // //w-writer, r-request
	HandleCoursePost(c *gin.Context)   // //w-writer, r-request
	HandleCoursePut(c *gin.Context)    // //w-writer, r-request
	HandleCourseDelete(c *gin.Context) // //w-writer, r-request
}

type courseHandler struct {
	courseService services.CourseService
}

func NewCourseHandler(courseService services.CourseService) CourseHandler {
	return &courseHandler{courseService: courseService}
}

// HandleCourseGet godoc
//
//	@Summary		Get course
//	@Description	Get course by id or all courses
//	@Tags			courses
//	@Produce		json
//	@Param			id	query		int	false	"Course ID"
//	@Success		200	{array}		entity.Course
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/courses [get]
func (courseHandler *courseHandler) HandleCourseGet(c *gin.Context) {
	var idStr string = c.Query("id")

	// GET ALL
	if idStr == "" {

		logger.Log.Info("Getting all courses")

		courses, err := courseHandler.courseService.GetAllCourses()
		if err != nil {
			c.Error(err)
			return
		}

		logger.Log.Debugf("Courses count: %d", len(courses))

		c.JSON(http.StatusOK, courses)
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
	logger.Log.Info("Getting course by id")

	logger.Log.Debugf("Course ID: %d", id)

	course, err := courseHandler.courseService.GetCourse(id)
	if err != nil {
		c.Error(err)
		return
	}

	logger.Log.Debugf("Course data: %+v", course)

	c.JSON(http.StatusOK, course)
}

// HandleCoursePost godoc
//
//	@Summary		Create course
//	@Description	Create new course
//	@Tags			courses
//	@Accept			json
//	@Produce		json
//	@Param			course	body		entity.Course	true	"Course"
//	@Success		201		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/courses [post]
func (courseHandler *courseHandler) HandleCoursePost(c *gin.Context) {
	var course entity.Course

	err := c.ShouldBindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	// INFO
	logger.Log.Info("Creating new course")

	// DEBUG
	logger.Log.Debugf("Course data: %+v", course)

	err = courseHandler.courseService.AddCourse(course)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "course created",
	})
}

// HandleCoursePut godoc
//
//	@Summary		Update course
//	@Description	Update existing course
//	@Tags			courses
//	@Accept			json
//	@Produce		json
//	@Param			course	body		entity.Course	true	"Updated Course"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/courses [put]
func (courseHandler *courseHandler) HandleCoursePut(c *gin.Context) {
	var updateCourse entity.Course
	err := c.ShouldBindJSON(&updateCourse)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}

	logger.Log.Info("Updating course")

	logger.Log.Debugf("Updated course data: %+v", updateCourse)

	err = courseHandler.courseService.UpdateCourse(updateCourse)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "course updated",
	})
}

// HandleCourseDelete godoc
//
//	@Summary		Delete course
//	@Description	Delete course by id
//	@Tags			courses
//	@Produce		json
//	@Param			id	query		int	true	"Course ID"
//	@Success		200	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/courses [delete]
func (courseHandler *courseHandler) HandleCourseDelete(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	logger.Log.Info("Deleting course")

	logger.Log.Debugf("Deleting course ID: %d", id)

	err = courseHandler.courseService.DeleteCourse(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "course deleted",
	})
}
