package db

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"remind/server/models"
)

var DB *gorm.DB

// ConnectDB initializes the database connection using GORM
func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables.")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully!")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("Database connection closed")
		}
	}
}

// InitSchema creates the database tables using GORM AutoMigrate
func InitSchema() error {
	if DB == nil {
		return fmt.Errorf("database not connected")
	}

	// AutoMigrate will create tables and update schema as needed
	err := DB.AutoMigrate(&models.Reminder{}, &models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema: %v", err)
	}

	log.Println("Database schema initialized successfully")
	return nil
}

