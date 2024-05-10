package routes

import (
	"net/http"
	"restAPI/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Short password or user with this login already exist"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "User created"})
}
