package routes

import (
	"example/goUdemyRest/models"
	"example/goUdemyRest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not sign user up"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "signed up successfully"})

}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "wrong user credentials", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wrong user credentials", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user authorized", "token": token})

}
