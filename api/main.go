package main

import (
	"golang-server/database"
	"golang-server/middleware"
	"golang-server/routes"
	"log"

	"github.com/gin-gonic/gin" // If using Gin
)

func main() {
	// connecting to database
	database.ConnectDatabase()
	router := gin.Default()

	// for local debugging purposes
	router.Use(middleware.CorsEnablement())

	// register routes and middleware
	routes.RegisterRoutes(router)

	log.Fatal(router.Run(":8080")) // Start the server on port 8080
}
