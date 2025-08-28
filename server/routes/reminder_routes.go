package routes


import (
	"github.com/gin-gonic/gin"
	"remind/server/controllers"
)

func ReminderRoutes(r *gin.Engine) {
	reminderGroup := r.Group("/reminders")
	{
	reminderGroup.GET("/reminders", controllers.GetReminders)
	}
}