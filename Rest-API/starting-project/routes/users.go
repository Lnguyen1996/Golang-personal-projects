package routes

import (
	"net/http"

	"example.com/rest-api/models"
	utils "example.com/rest-api/utiles"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Save user"})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})

		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if(err != nil){
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}
 
	context.JSON(http.StatusOK, gin.H{"message": "Login successfully","token":token})

}
