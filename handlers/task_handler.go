package handlers

import (
	"2_TaskManager/models"
	"2_TaskManager/store"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = store.NextID
	store.NextID++

	store.Tasks = append(store.Tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, store.Tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range store.Tasks {
		if id == fmt.Sprint(task.ID) {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range store.Tasks {
		if id == fmt.Sprint(task.ID) {
			updatedTask.ID = task.ID
			store.Tasks[i] = updatedTask

			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range store.Tasks {
		if id == fmt.Sprint(task.ID) {

			store.Tasks = append(store.Tasks[:i], store.Tasks[i+1:]...)

			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
