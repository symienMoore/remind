package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"remind/server/db"
	"remind/server/models"
)

func GetReminders(c *gin.Context) {
	var reminders []models.Reminder
	result := db.DB.Find(&reminders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	c.JSON(http.StatusOK, reminders)
}

func GetReminderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var reminder models.Reminder
	result := db.DB.First(&reminder, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func CreateReminder(c *gin.Context) {
	var reminder models.Reminder
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&reminder)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, reminder)
}

func UpdateReminder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var reminder models.Reminder
	result := db.DB.First(&reminder, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&reminder)
	c.JSON(http.StatusOK, reminder)
}

func DeleteReminder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var reminder models.Reminder
	result := db.DB.First(&reminder, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reminder not found"})
		return
	}

	db.DB.Delete(&reminder)
	c.JSON(http.StatusOK, gin.H{"message": "Reminder deleted"})
}

func SearchReminders(c *gin.Context) {
	searchTerm := c.Query("q")
	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search term 'q' is required"})
		return
	}

	var reminders []models.Reminder
	searchPattern := "%" + searchTerm + "%"
	
	result := db.DB.Where("title ILIKE ? OR description ILIKE ?", searchPattern, searchPattern).Find(&reminders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func GetDatabaseStats(c *gin.Context) {
	var count int64
	db.DB.Model(&models.Reminder{}).Count(&count)
	
	c.JSON(http.StatusOK, gin.H{
		"total_reminders": count,
		"database_url":    "***",
	})
}