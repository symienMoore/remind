package models

import (
	"time"
	"gorm.io/gorm"
)

type Reminder struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string         `json:"title" gorm:"size:255;not null"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// CreateReminderRequest represents the request payload for creating a reminder
type CreateReminderRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// UpdateReminderRequest represents the request payload for updating a reminder
type UpdateReminderRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// ReminderResponse represents the response payload for reminder operations
type ReminderResponse struct {
	Success bool      `json:"success"`
	Data    *Reminder `json:"data,omitempty"`
	Error   string    `json:"error,omitempty"`
}

// RemindersResponse represents the response payload for multiple reminders
type RemindersResponse struct {
	Success bool        `json:"success"`
	Data    []Reminder `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count"`
}
