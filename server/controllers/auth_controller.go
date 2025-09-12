package controllers

import (
	"github.com/gin-gonic/gin"
	"remind/server/auth"
	"remind/server/models"
	"golang.org/x/crypto/bcrypt"
	"remind/server/db"
)


func LogInUser(c *gin.Context) {

}



func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Store user in database
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return 
	}

	c.JSON(201, gin.H{
		"message": "User registered successfully",
		"token":   token,
		"user":    user,
	})
}