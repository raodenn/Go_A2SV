package models

type Task struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
	DueDate     string `json:"dueDate" bson:"dueDate" validate:"required,datetime=2006-01-02"`
	Status      string `json:"status" bson:"status" validate:"required,oneof=pending incomplete complete"`
}
