package main

import (
	"golang-server/routes"
	"log"

	"github.com/gin-gonic/gin" // If using Gin
)

func main() {
	router := gin.Default() // If using Gin

	// for local debugging purposes
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// register routes and middleware
	routes.RegisterRoutes(router)

	log.Fatal(router.Run(":8080")) // Start the server on port 8080
}
