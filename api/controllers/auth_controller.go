package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"golang-server/database"
	"golang-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handle login control
func LoginHandler(c *gin.Context) {
	var input utils.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password for comparison
	hashPaswrd := sha256.Sum256([]byte(input.Password))

	// try and find the user in the database
	userIds, err := database.FindUserWithCredentials(input.Username, hex.EncodeToString(hashPaswrd[:]))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// add session token to database
	err = database.CreateSession(userIds.Sessid, userIds.Id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// generate a jwt
	jwtToken, err := utils.GenerateJWT(userIds.Id, c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	utils.GenerateCookie(jwtToken, "token", c)
	utils.GenerateCookie(userIds.Sessid, "session_token", c)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

// registration handle conotrol
func RegisterHandler(c *gin.Context) {
	var input utils.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password for storage
	hashPaswrd := sha256.Sum256([]byte(input.Password))

	// check for existing credentials
	existing, err := database.FindUserWithUsername(input.Username)

	if existing || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Existing user has email registered"})
		return
	}

	// attempt to register the user in the database
	regUser, sessid, err := database.AddNewUser(input.Username, hex.EncodeToString(hashPaswrd[:]))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error registering user"})
		return
	}

	// add session token to database with id
	err = database.CreateSession(sessid, regUser)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// generate JWT
	jwtToken, err := utils.GenerateJWT(regUser, c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	// cookie generation and notification
	utils.GenerateCookie(jwtToken, "token", c)
	utils.GenerateCookie(sessid, "session_token", c)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}
