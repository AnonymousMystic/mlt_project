package main

import (
	"golang-server/database"
	"golang-server/middleware"
	"golang-server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connecting to database
	database.ConnectDatabase()
	router := gin.Default()

	// for local debugging purposes
	router.Use(middleware.CorsEnablement())

	// register routes and running JWT middleware
	routes.RegisterRoutes(router)

	// Start the server on port 8080
	log.Fatal(router.Run(":8080"))
}
