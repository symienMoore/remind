package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config/config"
	"server/routes/routes"
)

func main() {

	config.loadEnv()
	r := gin.Default()
	routes.RegisterRoutes(r)
	// Import your controllers and set up routes here.
	// Example:
	// userController := controllers.NewUserController()
	// router.GET("/users", userController.GetUsers)
	// router.POST("/users", userController.CreateUser)
	// router.GET("/users/:id", userController.GetUserByID)
	// router.PUT("/users/:id", userController.UpdateUser)
	// router.DELETE("/users/:id", userController.DeleteUser)

	// You can add more routes for other resources as needed.
	// e.g. router.GET("/posts", postController.GetPosts)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}