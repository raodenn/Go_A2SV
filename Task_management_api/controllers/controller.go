package controllers

import (
	"net/http"

	"Task_management_api/data"
	"Task_management_api/models"

	"github.com/gin-gonic/gin"
)

// Get all tasks
func GetTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Get one task by its ID
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, found, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
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

	updatedTask, found, changed, err := data.UpdateTaskByID(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
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

	if err := data.CreateTask(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
	}
	c.JSON(http.StatusCreated, newTask)
}

// Delete a task by ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	found, err := data.DeleteTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	if found {
		c.JSON(http.StatusOK, gin.H{"success": "Task deleted"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func RegisterUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, "Registered successfully")
}

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := data.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, foundUser)
}
