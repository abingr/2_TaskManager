package main

import (
	"2_TaskManager/db"
	"2_TaskManager/handlers"
	"2_TaskManager/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	db.ConnectDB()

	r := gin.Default()
	/* Endpoints */
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/tasks", handlers.GetTasks)
	protected.POST("/tasks", handlers.CreateTask)

	//r.POST("/tasks", handlers.CreateTask)       // Create
	//r.GET("/tasks", handlers.GetTasks)          // Read
	//r.GET("/tasks/:id", handlers.GetTaskByID)   // Read
	//r.PUT("/tasks/:id", handlers.UpdateTask)    // Update
	//r.DELETE("/tasks/:id", handlers.DeleteTask) // Delete
	r.Run(":8080")
}
