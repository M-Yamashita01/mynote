package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
)

type LogInParam struct {
	Email    string
	Password string
}

func PostLogin(c *gin.Context) {
	var logInParam LogInParam
	c.BindJSON(&logInParam)

	userProfile, err := model.FindUserProfile(logInParam.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed."})
		return
	}

	userId := userProfile.ID
	if !model.CorrectPassword(logInParam.Password, int(userId)) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed."})
		return
	}

	userToken, err := model.FindUserToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": userToken.Token})
}
