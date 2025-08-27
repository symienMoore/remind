package controllers

import "github.com/gin-gonic/gin"

func GetReminders(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get reminders endpoint"})
}