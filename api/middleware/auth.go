package middleware

import (
	"net/http"
	"time"

	"golang-server/database"
	"golang-server/utils"

	"github.com/gin-gonic/gin"
)

// ensures JWT is still valid before granting access to other resources
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		_, claims, err := utils.ValidateToken(token)

		// if JWT expired, check if there is a valid session
		if err != nil {
			// extract session id from cookie
			session_token, err := c.Cookie("session_token")

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			}

			// extract session start time from database
			dateTime, err := database.RetrieveSession(session_token)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			}

			// validate session life time
			if time.Since(dateTime) > 24*time.Hour {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Next()
	}
}
