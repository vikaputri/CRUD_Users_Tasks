package controllers

import (
	"fmt"
	"net/http"

	"github.com/vikaputri/CRUD_Users_Tasks/auth"
	"github.com/vikaputri/CRUD_Users_Tasks/database"
	"github.com/vikaputri/CRUD_Users_Tasks/models"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest

	db := database.GetDB()
	user := models.User{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	err := db.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Data with emails %v not found", request.Email),
		})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
