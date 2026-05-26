package middleware

import (
	"errors"
	"net/http"
	"studyProject/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		err := c.Errors.Last()

		if err == nil {
			return
		}

		// ERROR LOG
		logger.Log.Error(err.Err)

		// 404
		if errors.Is(err.Err, gorm.ErrRecordNotFound) {

			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "resource not found",
			})

			return
		}

		// 500
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
}
