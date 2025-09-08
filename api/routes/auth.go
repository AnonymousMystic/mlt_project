package routes

import (
	"golang-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(group *gin.RouterGroup) {
	auth := group.Group("/auth")

	auth.POST("/login", controllers.LoginHandler)
	auth.POST("/logout", controllers.LogoutHandler)
	auth.POST("/register", controllers.RegisterHandler)
}
