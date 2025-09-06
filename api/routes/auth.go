package routes

import (
	"golang-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(group *gin.RouterGroup) {
	auth := group.Group("/auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.RegisterHandler)
}
