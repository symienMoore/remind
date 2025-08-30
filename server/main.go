package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"context"
	"remind/server/config"
	"remind/server/routes"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	config.LoadEnv()
	err := godotenv.Load()
	if err != nil{
		log.Println("no .env file found")
	}
	
	connectString := os.Getenv("DATABASE_URL")
	if connectString == "" {
		fmt.Fprintf(os.Stderr, "DATABASE_URL not set\n")
		os.Exit(1)
	}

	ctx := context.Background()
	// Connect to the database
	conn, err := pgx.Connect(ctx, connectString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	fmt.Println("Connection established!!!!")

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()
	
	// Add CORS middleware for cross-origin requests
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})
	
	// API routes
	routes.ReminderRoutes(r)
	
	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"service": "remind-api",
			"status": "healthy",
		})
	})
	
	// API info endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "REmind API Server",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health": "/ping",
				"reminders": "/reminders",
			},
		})
	})

	// Get port from environment or use default
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	fmt.Printf("🚀 Starting REmind API server on port %s...\n", port)
	fmt.Printf("📱 Health check: http://localhost:%s/ping\n", port)
	fmt.Printf("🔗 API endpoints: http://localhost:%s/reminders\n", port)
	fmt.Println("🌐 CORS enabled for cross-origin requests")
	
	
	// Start the server
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("❌ Failed to run server: %v\n", err)
		os.Exit(1)
	}
}