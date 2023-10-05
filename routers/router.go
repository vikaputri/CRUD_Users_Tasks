package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vikaputri/CRUD_Users_Tasks/controllers"
	"github.com/vikaputri/CRUD_Users_Tasks/middlewares"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/token", controllers.GenerateToken)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	r := router.Group("/tasks")
	r.Use(middlewares.Auth())
	r.POST("/", controllers.CreateTask)
	r.GET("/", controllers.GetAllTasks)
	r.GET("/:id", controllers.GetTaskID)
	r.PUT("/:id", controllers.UpdateTask)
	r.DELETE("/:id", controllers.DeleteTask)

	return router

}
