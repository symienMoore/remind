package routes

import (
	"remind/server/controllers"
	"github.com/gin-gonic/gin"
)

func ReminderRoutes(r *gin.Engine) {
	reminderGroup := r.Group("/reminders")
	{
		reminderGroup.GET("/", controllers.GetReminders)
		reminderGroup.GET("/search", controllers.SearchReminders)
		reminderGroup.GET("/:id", controllers.GetReminderByID)
		reminderGroup.POST("/", controllers.CreateReminder)
		reminderGroup.PUT("/:id", controllers.UpdateReminder)
		reminderGroup.DELETE("/:id", controllers.DeleteReminder)
	}

	UserGroup := r.Group("/user")
	{
		UserGroup.GET("/")
	}

	AuthGroup := r.Group("/auth")
	{
		AuthGroup.POST("/signup", controllers.RegisterUser)
	}

	// Database stats endpoint
	r.GET("/db/stats", controllers.GetDatabaseStats)
}
