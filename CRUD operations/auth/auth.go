package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my-secret-key")

// generateToken generates JWT token for the given userID
func GenerateToken(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    fmt.Sprint(userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Authenticate is a middleware to authenticate requests
func Authenticate(c *gin.Context){
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid{
		c.JSON(401, gin.H{"error" : "Invalid or expired token"})
		c.Abort()
		return
	}
	
	c.Next()

}
