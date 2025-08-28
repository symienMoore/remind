package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"remind/server/config"
	"remind/server/routes"
)

func main() {
	config.LoadEnv()
	
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

	fmt.Println("🚀 Starting REmind API server on port 8080...")
	fmt.Println("📱 Health check: http://localhost:8080/ping")
	fmt.Println("🔗 API endpoints: http://localhost:8080/reminders")
	fmt.Println("🌐 CORS enabled for cross-origin requests")
	
	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("❌ Failed to run server: %v\n", err)
	}
}