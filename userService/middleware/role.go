package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		rolesInterface, _ := c.Get("roles")

		roles := rolesInterface.([]string)

		for _, r := range roles {

			if r == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
	}
}
