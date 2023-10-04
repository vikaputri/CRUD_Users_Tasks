package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vikaputri/CRUD_Users_Tasks/database"
	"github.com/vikaputri/CRUD_Users_Tasks/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(ctx *gin.Context) {
	db := database.GetDB()
	task := models.Task{}

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&task).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"message": "Failed to Create Task",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Task Created Successfully",
		"result":  task,
	})

}

func GetAllTasks(ctx *gin.Context) {
	db := database.GetDB()
	tasks := []models.Task{}
	err := db.Find(&tasks).Error
	if err != nil {
		panic(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"result":  tasks,
	})
}

func GetTaskID(ctx *gin.Context) {
	db := database.GetDB()
	tasks := []models.Task{}
	taskID := ctx.Param("id")
	err := db.Where("id = ?", taskID).First(&tasks).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("User with id %v not found", taskID),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  tasks,
		})
	}

}

func UpdateTask(ctx *gin.Context) {
	db := database.GetDB()
	task := models.Task{}
	task_id := ctx.Param("id")
	tasks := []models.Task{}

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(task_id)
	task.ID = uint(temp)

	err := db.Where("id = ?", task_id).First(&tasks).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Task with id %v not found", task_id),
		})
	} else {
		db.Model(&task).Updates(&task)
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"result":  task,
		})
	}

}

func DeleteTask(ctx *gin.Context) {
	db := database.GetDB()
	tasks := []models.Task{}
	taskID := ctx.Param("id")
	err := db.Where("id = ?", taskID).First(&tasks).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Task with id %v not found", taskID),
		})
	} else {
		db.Delete(&tasks)
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": true,
			"message": fmt.Sprintf("Task with id %v has been successfully delete", taskID),
		})
	}

}
