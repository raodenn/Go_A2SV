package router

import (
	"Task_management_api/controllers"
	"Task_management_api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.Login)

	authorized := router.Group("/tasks")
	authorized.Use(middleware.Authenticate())
	{
		authorized.GET("/", controllers.GetTasks)
		authorized.GET("/:id", controllers.GetTaskById)
		authorized.PUT("/:id", controllers.UpdateTask)
		authorized.POST("/", middleware.AuthorizeRoles("ADMIN", "USER"), controllers.CreateTask)
		authorized.DELETE("/:id", middleware.AuthorizeRoles("ADMIN"), controllers.DeleteTask)

	}
	return router
}
