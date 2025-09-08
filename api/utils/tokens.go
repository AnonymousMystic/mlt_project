package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// generates a secret
func generateSecret() {
	var err error

	e := godotenv.Load() // Load .env file
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	key := make([]byte, 32)
	_, err = rand.Read(key)

	for err != nil {
		fmt.Println("Error:", err)
		_, err = rand.Read(key)
	}

	os.Setenv("JWT_SECRET", base64.StdEncoding.EncodeToString(key))
}

// Create token
func GenerateToken(userID uint) (string, error) {
	// generate a secret if none exists
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
		"exp":     time.Now().Add(time.Hour).Unix(), // 1-day expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// Validate JWT and return claims
func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	jwtSecret, e := base64.StdEncoding.DecodeString(os.Getenv("JWT_SECRET"))

	if e != nil {
		return nil, nil, e
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, err
}

// Generate cookies for tokens
func GenerateCookie(tokenString string, name string, c *gin.Context) {
	c.SetCookie(
		name,
		tokenString,
		360,
		"/",
		"",
		false,
		true,
	)

	c.Header("Set-Cookie", name+"="+tokenString+"; Path=/; HttpOnly; Secure; SameSite=None")
}

// Generates input for token generation
func GenerateTokenInput(stringInput string, c *gin.Context) uint {
	tokenInput := StringToNumber(stringInput)
	fmt.Println(tokenInput)

	return uint(tokenInput)
}

// convert id to valid number
func StringToNumber(userId string) uint {
	hash := sha256.Sum256([]byte(userId))
	// Take the first 8 bytes and convert to uint64
	return uint(binary.BigEndian.Uint64(hash[:8]))
}
