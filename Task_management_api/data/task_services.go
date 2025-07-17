package data

import (
	"Task_management_api/models"
	"fmt"
)

// Dummy tasks
var tasks = []models.Task{
	{
		ID:          "1",
		Title:       "Buy groceries",
		Description: "Milk, Bread, Eggs",
		DueDate:     "2025-07-20",
		Status:      "pending",
	},
	{
		ID:          "2",
		Title:       "Finish report",
		Description: "Complete the quarterly report",
		DueDate:     "2025-07-25",
		Status:      "incomplete",
	},
	{
		ID:          "3",
		Title:       "Clean the house",
		Description: "Vacuum and mop floors",
		DueDate:     "2025-07-22",
		Status:      "complete",
	},
}

// To keep track of the next ID
var nextId = 4

// Return all tasks
func GetAllTasks() []models.Task {
	return tasks
}

// Find a task by its ID
func GetTaskByID(id string) (models.Task, bool) {
	for _, t := range tasks {
		if t.ID == id {
			return t, true
		}
	}
	return models.Task{}, false
}

// Update a task if it exists and has changes
func UpdateTaskByID(id string, newTask models.Task) (models.Task, bool, bool) {
	for i, t := range tasks {
		// Check if ID matches
		if t.ID == id {
			// If no changes, return false
			if t.Title == newTask.Title &&
				t.Description == newTask.Description &&
				t.DueDate == newTask.DueDate &&
				t.Status == newTask.Status {
				return t, true, false
			}
			// Update task
			newTask.ID = id
			tasks[i] = newTask
			return newTask, true, true
		}
	}
	return models.Task{}, false, false
}

// Add a new task
func CreateTask(task *models.Task) {
	task.ID = fmt.Sprintf("%d", nextId)
	nextId++
	tasks = append(tasks, *task)
}

// Remove a task by its ID
func DeleteTaskById(id string) bool {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}
