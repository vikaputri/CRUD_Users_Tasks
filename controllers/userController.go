package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vikaputri/CRUD_Users_Tasks/database"
	"github.com/vikaputri/CRUD_Users_Tasks/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"message": "Failed to Create User",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Created Successfully",
		"result":  user,
	})

}

func GetAllUsers(ctx *gin.Context) {
	db := database.GetDB()
	users := []models.User{}
	err := db.Find(&users).Error
	if err != nil {
		panic(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  users,
	})
}

func GetUserID(ctx *gin.Context) {
	db := database.GetDB()
	users := []models.User{}
	userID := ctx.Param("id")
	err := db.Where("id = ?", userID).First(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("User with id %v not found", userID),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  users,
		})
	}

}

func UpdateUser(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}
	user_id := ctx.Param("id")
	users := []models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(user_id)
	user.ID = uint(temp)

	err := db.Where("id = ?", user_id).First(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("User with id %v not found", user_id),
		})
	} else {
		db.Model(&user).Updates(&user)
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"result":  user,
		})
	}

}

func DeleteUser(ctx *gin.Context) {
	db := database.GetDB()
	users := []models.User{}
	userID := ctx.Param("id")
	err := db.Where("id = ?", userID).First(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("User with id %v not found", userID),
		})
	} else {
		db.Delete(&users)
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": true,
			"message": fmt.Sprintf("User with id %v has been successfully delete", userID),
		})
	}

}
