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
	
	// Add middleware for debugging
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	// API routes should come first
	routes.ReminderRoutes(r)
	
	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		  "message": "pong",
		})
	})
	  

    // Import your controllers and set up routes here.
    // Example:
    // userController := controllers.NewUserController()
    // router.GET("/users", userController.GetUsers)
    // router.POST("/users", userController.CreateUser)
    // router.GET("/users/:id", userController.GetUserByID)
    // router.PUT("/users/:id", userController.UpdateUser)
    // router.DELETE("/users/:id", userController.UpdateUser)

    // You can add more routes for other resources as needed.
    // e.g. router.GET("/posts", postController.GetPosts)
	// Serve static files from the Angular build (after API routes)
	r.Static("/assets", "./ui/assets")
	r.StaticFile("/", "./ui/index.html")
	r.StaticFile("/index.html", "./ui/index.html")
	
	// Catch-all route for Angular client-side routing (must be last)
	r.NoRoute(func(c *gin.Context) {
		c.File("./ui/index.html")
	})

	fmt.Println("🚀 Starting REmind server on port 8080...")
	fmt.Println("📱 Health check: http://localhost:8080/ping")
	fmt.Println("🔗 API endpoints: http://localhost:8080/reminders")
	
	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("❌ Failed to run server: %v\n", err)
	}
}