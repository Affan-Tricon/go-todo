package routes

import (
	"todo/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.POST("/users", controllers.CreateUsers)
		api.GET("/users/*id", controllers.GetUsers)
		api.DELETE("/users/:id", controllers.DeleteUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.POST("/login", controllers.Login)
	}
	{
		api.POST("/todo", controllers.CreateTodo)
		api.GET("/todo/*id", controllers.GetTodo)
		api.PUT("/todo/:id", controllers.UpdateTodo)
		api.DELETE("/todo/:id", controllers.DeleteTodo)
	}

}
