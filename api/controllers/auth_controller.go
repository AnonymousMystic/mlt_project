package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"golang-server/database"
	"golang-server/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// handle login control
func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password
	hashPaswrd := sha256.Sum256([]byte(input.Password))

	// try and find the user in the database
	authUser, err := database.FindUserWithCredentials(input.Username, hex.EncodeToString(hashPaswrd[:]))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenInput, err := strconv.Atoi(authUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UUID error"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(uint(tokenInput))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	utils.GenerateCookie(token, c)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

func RegisterHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password for storage
	hashPaswrd := sha256.Sum256([]byte(input.Password))

	// try and find the user in the database
	regUser, err := database.AddNewUser(input.Username, hex.EncodeToString(hashPaswrd[:]))

	print(regUser)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error registering user"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(1) // Simulate userID = 1
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	// cookie generation and notification
	utils.GenerateCookie(token, c)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}
