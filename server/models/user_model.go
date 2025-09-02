package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID            uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string     `json:"username" gorm:"size:100;uniqueIndex;unique;not null"`
	Email         string     `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Password      string     `json:"password" gorm:"size:255;not null;unique"`
	ProfilePic    string     `json:"profilePic" gorm:"size:500"`
	IsActive      bool       `json:"isActive" gorm:"default:true"`
	CreatedAt     time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeactivatedAt *time.Time `json:"deactivatedAt,omitempty" gorm:"index"`
}
