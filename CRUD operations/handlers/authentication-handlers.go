package handlers

import (
	"go-gin-postgres/auth"
	"go-gin-postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login authenticates user credentials and generates JWT token
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		// Bind request body to User struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if required fields are provided
		if user.Name == "" || user.Email == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name, Email, and Password are required fields"})
			return
		}

		// Generate JWT token
		tokenString, err := auth.GenerateToken(1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		// Return token in response
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}