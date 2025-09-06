package routes

import (
	"github.com/gin-gonic/gin"

	"golang-server/controllers"
)

func RegisterUserRoutes(group *gin.RouterGroup) {
	user := group.Group("/user")

	user.GET("/profile", controllers.Profile)
}
