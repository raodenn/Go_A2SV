package data

import (
	"Task_management_api/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Return all tasks
func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := TaskCollections.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var res []models.Task
	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// Find a task by its ID
func GetTaskByID(id string) (models.Task, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var task models.Task
	filter := bson.M{"_id": id}
	err := TaskCollections.FindOne(ctx, filter).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, false, nil
		}
		return models.Task{}, false, err
	}
	return task, true, nil
}

// Update a task if it exists and has changes
func UpdateTaskByID(id string, newTask models.Task) (models.Task, bool, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}

	var existing models.Task
	err := TaskCollections.FindOne(ctx, filter).Decode(&existing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, false, false, nil
		}
		return models.Task{}, false, false, err
	}
	if existing.Title == newTask.Title &&
		existing.Description == newTask.Description &&
		existing.DueDate == newTask.DueDate &&
		existing.Status == newTask.Status {
		return existing, true, false, nil
	}
	newTask.ID = id
	_, err = TaskCollections.ReplaceOne(ctx, filter, newTask)
	if err != nil {
		return models.Task{}, false, false, err
	}

	return newTask, true, true, nil
}

// Add a new task
func CreateTask(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	task.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	_, err := TaskCollections.InsertOne(ctx, task)
	return err
}

// Remove a task by its ID
func DeleteTaskById(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := TaskCollections.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}
	return res.DeletedCount > 0, nil
}
