package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing token",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, _, err := new(jwt.Parser).ParseUnverified(
			tokenString,
			jwt.MapClaims{},
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid claims",
			})
			return
		}

		username, _ := claims["preferred_username"].(string)

		var roles []string

		if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {

			if rawRoles, ok := realmAccess["roles"].([]interface{}); ok {

				for _, role := range rawRoles {

					if roleStr, ok := role.(string); ok {
						roles = append(roles, roleStr)
					}
				}
			}
		}

		c.Set("username", username)
		c.Set("roles", roles)

		c.Next()
	}
}
