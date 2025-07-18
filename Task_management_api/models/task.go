package models

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" binding:"required"`
	Status      string `json:"status" binding:"required,oneof= complete pending incomplete"`
}
