package routes

import (
	"example/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup user"})

		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signup successful"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = user.ValidateCredentails()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "auth failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}
