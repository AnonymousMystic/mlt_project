package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey")

// generates a secret
func generateSecret() {
	var err error

	key := make([]byte, 32)
	_, err = rand.Read(key)

	for err != nil {
		fmt.Println("Error:", err)
		_, err = rand.Read(key)
	}

	os.Setenv("JWT_SECRET", base64.StdEncoding.EncodeToString(key))
}

// Create JWT token
func GenerateToken(userID uint) (string, error) {
	// generate a JWT secret if none exists
	if len(os.Getenv("JWT_SECRET")) == 0 {
		generateSecret()
	}

	// extract jwtSecret
	jwtSecret, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1-day expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// Validate token and return claims
func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, err
}
