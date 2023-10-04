package routers

import (
	"github.com/vikaputri/CRUD_Users_Tasks/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTaskID)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	return router

}
