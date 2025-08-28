package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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
	
	// Static file serving for Angular app
	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./static"
	}

	// Serve assets and common static files if available
	r.Static("/assets", filepath.Join(staticDir, "assets"))
	r.StaticFile("/favicon.ico", filepath.Join(staticDir, "favicon.ico"))

	// SPA fallback to index.html, otherwise API info
	r.NoRoute(func(c *gin.Context) {
		indexPath := filepath.Join(staticDir, "index.html")
		if _, err := os.Stat(indexPath); err == nil {
			c.File(indexPath)
			return
		}
		c.JSON(http.StatusOK, gin.H{
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
	fmt.Printf("📦 Serving static files from: %s (if present)\n", staticDir)
	fmt.Println("🌐 CORS enabled for cross-origin requests")
	
	// Start the server
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("❌ Failed to run server: %v\n", err)
		os.Exit(1)
	}
}