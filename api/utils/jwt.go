package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GenerateJWT(tokenInput string, c *gin.Context) (string, error) {
	// convert session token into usable input for token
	jwtInput := GenerateTokenInput(tokenInput, c)

	// Generate token
	token, err := GenerateToken(jwtInput)
	if err != nil {
		return "", errors.New("token error")
	}

	return token, nil
}
