package routes

import (
	"github.com/gin-gonic/gin"
	"remind/server/controllers"
)

func ReminderRoutes(r *gin.Engine) {
	reminderGroup := r.Group("/reminders")
	{
		reminderGroup.GET("/", controllers.GetReminders)
		reminderGroup.GET("/:id", controllers.GetReminderByID)
		reminderGroup.POST("/", controllers.CreateReminder)
		reminderGroup.PUT("/:id", controllers.UpdateReminder)
		reminderGroup.DELETE("/:id", controllers.DeleteReminder)
	}
}