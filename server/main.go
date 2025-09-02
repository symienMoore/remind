package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"remind/server/config"
	"remind/server/db"
	"remind/server/routes"
)

func main() {
	config.LoadEnv()
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.CloseDB()

	if err := db.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// CORS
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
	//routes.UserRoutes(r)

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"service": "remind-api",
			"status":  "healthy",
		})
	})

	// API info endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "REmind API Server",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":    "/ping",
				"reminders": "/reminders",
				"search":    "/reminders/search?q=<search_term>",
				"db_stats":  "/db/stats",
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
	fmt.Println("🗄️  Database connected using GORM")

	// Start the server
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("❌ Failed to run server: %v\n", err)
		os.Exit(1)
	}
}
