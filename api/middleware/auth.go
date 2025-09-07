package middleware

import (
	"net/http"

	"golang-server/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")

		println(token)
		println(err)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		_, claims, err := utils.ValidateToken(token)

		println(err)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID in context
		println("user_id", uint(claims["user_id"].(float64)))
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Next()
	}
}
