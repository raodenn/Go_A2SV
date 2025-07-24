package usecases

import (
	"context"
	domain "task_manager/domain"
)

type TaskUseCase struct {
	Repo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		Repo: repo,
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task *domain.Task) error {
	return uc.Repo.CreateTask(ctx, task)
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	return uc.Repo.GetTask(ctx, id)
}

func (uc *TaskUseCase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return uc.Repo.GetAllTasks(ctx)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, task *domain.Task) error {
	return uc.Repo.UpdateTask(ctx, task)
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	return uc.Repo.DeleteTask(ctx, id)
}
