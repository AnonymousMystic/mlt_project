package routes

import (
	"golang-server/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		RegisterAuthRoutes(api)
		api.Use(middleware.JWTAuthMiddleware())
		RegisterUserRoutes(api)
	}
}
