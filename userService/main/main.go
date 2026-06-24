package main

import (
	"log"
	"userService/handlers"
	"userService/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user-service works",
		})
	})

	// AUTH
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/refresh", handlers.Refresh)

	// Только ADMIN может создавать пользователей
	admin := r.Group("/")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RequireRole("ROLE_ADMIN"),
	)

	admin.POST("/users", handlers.CreateUser)

	// ADMIN и TEACHER могут обновлять пользователя
	adminTeacher := r.Group("/")
	adminTeacher.Use(
		middleware.AuthMiddleware(),
	)

	adminTeacher.PUT("/users/:username", handlers.UpdateUser)
	adminTeacher.PUT("/users/:username/password", handlers.ChangePassword)

	log.Println("User Service started on :8082")

	err := r.Run(":8082")
	if err != nil {
		log.Fatal(err)
	}
}
