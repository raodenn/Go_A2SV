package repositories

import (
	"context"
	"errors"
	domain "task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	Collection *mongo.Collection
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	UserId      string             `bson:"user_id"`
	DueDate     string             `bson:"due_date"`
	Status      string             `bson:"status"`
}

func NewTaskRepo() domain.TaskRepository {
	col := initdb().Collection("tasks")
	return &TaskRepo{
		Collection: col,
	}
}

func (r *TaskRepo) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task id")
	}
	var taskModel Task
	err = r.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&taskModel)
	if err != nil {
		return nil, err
	}
	task := &domain.Task{
		ID:          taskModel.ID.Hex(),
		Title:       taskModel.Title,
		Description: taskModel.Description,
		DueDate:     taskModel.DueDate,
		Status:      taskModel.Status,
	}
	return task, nil
}

func (r *TaskRepo) CreateTask(ctx context.Context, task *domain.Task) error {
	doc := Task{
		ID:          primitive.NewObjectID(),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status,
	}
	_, err := r.Collection.InsertOne(ctx, doc)
	return err
}

func (r *TaskRepo) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []*domain.Task
	for cursor.Next(ctx) {
		var doc Task
		if err := cursor.Decode(&doc); err != nil {
			continue
		}
		res = append(res, &domain.Task{
			ID:          doc.ID.Hex(),
			Title:       doc.Title,
			Description: doc.Description,
			DueDate:     doc.DueDate,
			Status:      doc.Status,
		})
	}
	return res, nil
}

func (r *TaskRepo) DeleteTask(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task id")
	}
	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (r *TaskRepo) UpdateTask(ctx context.Context, task *domain.Task) error {
	objectId, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
		},
	}
	_, err = r.Collection.UpdateByID(ctx, objectId, update)
	return err
}
