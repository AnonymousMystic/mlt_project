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

		_, _, err = utils.ValidateToken(token)

		// if JWT expired, check if there is a valid session
		if err != nil {
			// extract session id from cookie
			session_token, err := c.Cookie("session_token")

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
				return
			}

			// extract session start time from database
			dateTime, err := database.RetrieveSession(session_token)

			if time.Time.Equal(dateTime, time.Time{}) || err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
				return
			}

			// validate session life time
			if time.Since(dateTime) > 24*time.Hour {
				err = database.RemoveAndInvalidateSession(session_token)

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "session could not be removed"})
					c.Abort()
				}

				c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
				return
			}

			// get uuid from database to generate new jwt
			uuid, err := database.FindUserFromSession(session_token)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}

			// reissue JWT
			jwtToken, err := utils.GenerateJWT(uuid, c)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				return
			}
			// cookie generation and notification
			utils.GenerateCookie(jwtToken, "token", 3600, c)
		}

		// Set user ID in context
		c.JSON(http.StatusOK, gin.H{"message": "session valid"})
		c.Next()
	}
}
