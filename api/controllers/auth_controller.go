package controllers

import (
	"fmt"
	"golang-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulated auth
	if input.Username != "user@example.com" || input.Password != "password123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(1) // Simulate userID = 1
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	fmt.Println(token)

	// securely prepare cookie
	c.SetCookie(
		"token", // cookie name
		token,   // token value
		3600,    // maxAge in seconds
		"/",     // path
		"",      // domain ("" means current domain)
		true,    // secure (HTTPS only)
		true,    // httpOnly (not accessible via JS)
	)

	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

func RegisterHandler(c *gin.Context) {
	// Dummy register logic
	c.JSON(http.StatusOK, gin.H{"message": "Register successful"})
}
