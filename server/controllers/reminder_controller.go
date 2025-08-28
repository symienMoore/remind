package controllers

import "github.com/gin-gonic/gin"

func GetReminders(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get reminders endpoint"})
}

func GetReminderByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get reminder by ID", "id": id})
}

func CreateReminder(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Create reminder endpoint"})
}

func UpdateReminder(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Update reminder", "id": id})
}

func DeleteReminder(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Delete reminder", "id": id})
}