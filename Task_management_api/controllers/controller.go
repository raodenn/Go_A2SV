package controllers

import (
	"net/http"

	"Task_management_api/data"
	"Task_management_api/models"

	"github.com/gin-gonic/gin"
)

// Get all tasks
func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// Get one task by its ID
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, found := data.GetTaskByID(id)
	if !found {
		// Task not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	// Task found
	c.JSON(http.StatusOK, gin.H{"task:": task})
}

// Update a task by ID
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var update models.Task

	// Check if the request body is valid
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, found, changed := data.UpdateTaskByID(id, update)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if !changed {
		c.JSON(http.StatusNotModified, gin.H{"message": "no changes detected"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// Create a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task

	// Check if the request body is valid
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.CreateTask(&newTask)
	c.JSON(http.StatusCreated, newTask)
}

// Delete a task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	found := data.DeleteTaskById(id)

	if found {
		c.JSON(http.StatusOK, gin.H{"success": "Task deleted"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
