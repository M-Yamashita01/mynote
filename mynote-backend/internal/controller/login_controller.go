package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
)

func ShowLoginPage() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	}
}

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

	userId := userProfile.Model.ID
	if !model.CorrectPassword(logInParam.Password, int(userId)) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed."})
		return
	}

	userToken, err := model.FindUserToken(int(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": userToken.Token})
}
