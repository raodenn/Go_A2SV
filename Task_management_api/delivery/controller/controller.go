package controller

import (
	"net/http"
	domain "task_manager/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type Tctrl struct {
	tuc domain.TaskUseCase
}
type Uctrl struct {
	uuc    domain.UserUseCase
	jwtSvc domain.JwtSvc
}

type UserDTO struct {
	ID        string    `json:"id,omitempty"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	UserType  string    `json:"usertype" binding:"required"`
	Token     *string   `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type TaskDTO struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	Status      string `json:"status,omitempty"`
}

func NewUserCtrl(uc domain.UserUseCase, jwt domain.JwtSvc) *Uctrl {
	return &Uctrl{
		uuc:    uc,
		jwtSvc: jwt,
	}
}

func NewTaskCtrl(tuc domain.TaskUseCase) *Tctrl {
	return &Tctrl{
		tuc: tuc,
	}
}

func (Uctrl *Uctrl) Register(c *gin.Context) {
	var input UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := &domain.User{
		Username: input.Username,
		Password: &input.Password,
		UserType: input.UserType,
	}
	ctx := c.Request.Context()
	err := Uctrl.uuc.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func (Uctrl *Uctrl) Login(c *gin.Context) {
	var input LoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := domain.User{
		Username: input.Username,
		Password: &input.Password,
	}
	ctx := c.Request.Context()
	token, err := Uctrl.uuc.Login(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (Tctrl *Tctrl) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	tasks, err := Tctrl.tuc.GetAllTasks(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (Tctrl *Tctrl) GetTask(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	task, err := Tctrl.tuc.GetTask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (Tctrl *Tctrl) CreateTask(c *gin.Context) {
	var newTask TaskDTO

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := &domain.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		DueDate:     newTask.DueDate,
		Status:      newTask.Status,
	}

	ctx := c.Request.Context()
	if err := Tctrl.tuc.CreateTask(ctx, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

func (Tctrl *Tctrl) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var update TaskDTO
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &domain.Task{
		ID:          id,
		Title:       update.Title,
		Description: update.Description,
		DueDate:     update.DueDate,
		Status:      update.Status,
	}

	ctx := c.Request.Context()
	err := Tctrl.tuc.UpdateTask(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

func (Tctrl *Tctrl) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	err := Tctrl.tuc.DeleteTask(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
