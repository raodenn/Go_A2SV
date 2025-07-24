package router

import (
	controller "task_manager/delivery/controller"
	"task_manager/infrastructure"
	"task_manager/repositories"
	usecases "task_manager/usecases"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	taskRepo := repositories.NewTaskRepo()
	userRepo := repositories.NewUserRepo()

	PassSvc := infrastructure.NewPasswordSvc()

	duration := time.Hour * 24
	jwtSvc := infrastructure.NewJWTService("cleanarch", duration)

	Uuc := usecases.NewUserUseCase(userRepo, PassSvc, jwtSvc)
	Tuc := usecases.NewTaskUseCase(taskRepo)

	Tctrl := controller.NewTaskCtrl(Tuc)
	Uctrl := controller.NewUserCtrl(Uuc, jwtSvc)

	authMiddleware := infrastructure.NewAuthMiddleware(jwtSvc)

	r.POST("/users/register", Uctrl.Register)
	r.POST("/users/login", Uctrl.Login)

	r.GET("/tasks", authMiddleware.AuthorizeJWT("user", "admin"), Tctrl.GetAllTasks)
	r.GET("/tasks/:id", authMiddleware.AuthorizeJWT("user", "admin"), Tctrl.GetTask)
	r.PUT("/tasks/:id", authMiddleware.AuthorizeJWT("user", "admin"), Tctrl.UpdateTask)
	r.POST("/tasks", authMiddleware.AuthorizeJWT("admin"), Tctrl.CreateTask)
	r.DELETE("/tasks/:id", authMiddleware.AuthorizeJWT("admin"), Tctrl.DeleteTask)

	r.Run(":8080")
}
