package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		rolesInterface, exists := c.Get("roles")

		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "roles not found",
			})
			return
		}

		roles, ok := rolesInterface.([]string)

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "invalid roles",
			})
			return
		}

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
