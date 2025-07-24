package domain

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        string
	Username  string
	Password  *string
	UserType  string
	Token     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
}

type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     string
	Status      string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, username string) (*User, error)
}

type TaskRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetTask(ctx context.Context, id string) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id string) error
}

type PasswordSvc interface {
	HashPassword(password string) *string
	VerifyPassword(pass string, foundpass string) (bool, error)
}

type JwtSvc interface {
	GenerateToken(userID string, userType string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user *User) error
	Login(ctx context.Context, user User) (string, error)
}

type TaskUseCase interface {
	CreateTask(ctx context.Context, task *Task) error
	GetTask(ctx context.Context, id string) (*Task, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id string) error
}
