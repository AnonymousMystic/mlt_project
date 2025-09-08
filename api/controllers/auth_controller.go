package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"golang-server/database"
	"golang-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// handle login control
func LoginHandler(c *gin.Context) {
	var input utils.LoginInput
	var sessid string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password for comparison
	hashPaswrd := sha256.Sum256([]byte(input.Password))

	// try and find the user in the database using their provided credentials
	userIds, err := database.FindUserWithCredentials(input.Username, hex.EncodeToString(hashPaswrd[:]))

	if len(userIds.Sessid) == 0 {
		sessid = uuid.New().String()
	} else {
		sessid = userIds.Sessid
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// add session token to database
	err = database.CreateSession(sessid, userIds.Id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// generate a jwt
	jwtToken, err := utils.GenerateJWT(userIds.Id, c)
	print(err == nil)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	utils.GenerateCookie(jwtToken, "token", 360, c)
	utils.GenerateCookie(userIds.Sessid, "session_token", 8640, c)
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
	utils.GenerateCookie(jwtToken, "token", 3600, c)
	utils.GenerateCookie(sessid, "session_token", 86400, c)
	c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
}

func LogoutHandler(c *gin.Context) {
	// extract session token
	sessId, err := c.Cookie("session_token")

	if len(sessId) == 0 && err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not extract session token"})
		return
	}

	// revoke session token
	err = database.RemoveAndInvalidateSession(sessId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "could not revoke session token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}
