package handlers

import (
	"2_TaskManager/db"
	"2_TaskManager/models"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id"

	err := db.Conn.QueryRow(context.Background(),
		query, task.Title, task.Completed).Scan(&task.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {

	log.Println("Fetching tasks...")

	rows, err := db.Conn.Query(context.Background(),
		"SELECT id, title, completed FROM tasks")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		log.Println("Row found")
		err := rows.Scan(&task.ID, &task.Title, &task.Completed)
		if err != nil {
			continue
		}

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, title, completed FROM tasks WHERE id=$1", id).Scan(&task.ID, &task.Title, &task.Completed)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
	UPDATE tasks
	SET title=$1, completed=$2
	WHERE id=$3
	`
	_, err := db.Conn.Exec(context.Background(),
		query, task.Title, task.Completed, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	task.ID, _ = strconv.Atoi(id)

	c.JSON(http.StatusOK, task)

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Conn.Exec(context.Background(),
		"DELETE FROM tasks WHERE id=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
