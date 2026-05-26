//	@title			Study Project API
//	@version		1.0
//	@description	REST API for courses, chapters and lessons
//	@host			localhost:8080
//	@BasePath		/

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"studyProject/handlers"
	"studyProject/logger"
	"studyProject/middleware"
	"studyProject/repositories"
	"studyProject/services"
)
import _ "github.com/lib/pq"
import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "studyProject/docs"
)

var gormDB *gorm.DB

func InitDB() {
	connection := "user=postgres password=admin dbname=goDB host=db port=5432 sslmode=disable"
	var err error
	gormDB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed")
	}
	//errTwo := gormDB.AutoMigrate(&entity.Course{}, &entity.Chapter{}, &entity.Lesson{})
	//if errTwo != nil {
	//	log.Fatal("Auto migration failed")
	//}
}

func CloseDB() {
	s, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	errTwo := s.Close()
	if errTwo != nil {
		log.Fatal("Closing failed")
	}
}

func main() {
	logger.InitLogger()
	InitDB()
	defer CloseDB()

	var courseRepository repositories.CourseRepository
	courseRepo := repositories.NewCourseRepository(gormDB)
	courseRepository = courseRepo

	var courseService services.CourseService
	courseServ := services.NewCourseService(courseRepository)
	courseService = courseServ

	var courseHandler handlers.CourseHandler
	courseHand := handlers.NewCourseHandler(courseService)
	courseHandler = courseHand

	var chapterRepository repositories.ChapterRepository
	chapterRepo := repositories.NewChapterRepository(gormDB)
	chapterRepository = chapterRepo

	var chapterService services.ChapterService
	chapterServ := services.NewChapterService(chapterRepository)
	chapterService = chapterServ

	var chapterHandler handlers.ChapterHandler
	chapterHand := handlers.NewChapterHandler(chapterService)
	chapterHandler = chapterHand

	var lessonRepository repositories.LessonRepository
	lessonRepo := repositories.NewLessonRepository(gormDB)
	lessonRepository = lessonRepo

	var lessonService services.LessonService
	lessonServ := services.NewLessonService(lessonRepository)
	lessonService = lessonServ

	var lessonHandler handlers.LessonHandler
	lessonHand := handlers.NewLessonHandler(lessonService)
	lessonHandler = lessonHand

	// gin router
	r := gin.Default()

	//Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.ErrorMiddleware())

	// COURSES

	r.GET("/courses", courseHandler.HandleCourseGet)

	r.POST("/courses", courseHandler.HandleCoursePost)

	r.PUT("/courses", courseHandler.HandleCoursePut)

	r.DELETE("/courses", courseHandler.HandleCourseDelete)

	// CHAPTERS

	r.GET("/chapters", chapterHandler.HandleChapterGet)

	r.POST("/chapters", chapterHandler.HandleChapterPost)

	r.PUT("/chapters", chapterHandler.HandleChapterPut)

	r.DELETE("/chapters", chapterHandler.HandleChapterDelete)

	// LESSONS

	r.GET("/lessons", lessonHandler.HandleLessonGet)

	r.POST("/lessons", lessonHandler.HandleLessonPost)

	r.PUT("/lessons", lessonHandler.HandleLessonPut)

	r.DELETE("/lessons", lessonHandler.HandleLessonDelete)

	logger.Log.Info("Server started on :8080")

	err := r.Run(":8080")

	if err != nil {
		logger.Log.Error(err)
	}
}
