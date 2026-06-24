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

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middleware
	r.Use(middleware.ErrorMiddleware())

	// Все защищенные маршруты
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// ==================== COURSES ====================

	authorized.GET(
		"/courses",
		courseHandler.HandleCourseGet,
	)

	authorized.POST(
		"/courses",
		middleware.RequireRole("ROLE_ADMIN"),
		courseHandler.HandleCoursePost,
	)

	authorized.PUT(
		"/courses",
		middleware.RequireRole("ROLE_ADMIN"),
		courseHandler.HandleCoursePut,
	)

	authorized.DELETE(
		"/courses",
		middleware.RequireRole("ROLE_ADMIN"),
		courseHandler.HandleCourseDelete,
	)

	// ==================== CHAPTERS ====================

	authorized.GET(
		"/chapters",
		chapterHandler.HandleChapterGet,
	)

	authorized.POST(
		"/chapters",
		middleware.RequireRole("ROLE_ADMIN"),
		chapterHandler.HandleChapterPost,
	)

	authorized.PUT(
		"/chapters",
		middleware.RequireRole("ROLE_ADMIN"),
		chapterHandler.HandleChapterPut,
	)

	authorized.DELETE(
		"/chapters",
		middleware.RequireRole("ROLE_ADMIN"),
		chapterHandler.HandleChapterDelete,
	)

	// ==================== LESSONS ====================

	authorized.GET(
		"/lessons",
		lessonHandler.HandleLessonGet,
	)

	authorized.POST(
		"/lessons",
		middleware.RequireRole("ROLE_ADMIN"),
		lessonHandler.HandleLessonPost,
	)

	authorized.PUT(
		"/lessons",
		middleware.RequireRole("ROLE_ADMIN"),
		lessonHandler.HandleLessonPut,
	)

	authorized.DELETE(
		"/lessons",
		middleware.RequireRole("ROLE_ADMIN"),
		lessonHandler.HandleLessonDelete,
	)

	logger.Log.Info("Server started on :8080")

	err := r.Run(":8080")
	if err != nil {
		logger.Log.Error(err)
	}
}
