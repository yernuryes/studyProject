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

		claims := token.Claims.(jwt.MapClaims)

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

		c.Set("roles", roles)

		c.Next()
	}
}
