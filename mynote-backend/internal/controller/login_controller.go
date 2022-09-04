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

func PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	userProfile, err := model.FindUserProfile(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed."})
		return
	}

	userId := userProfile.Model.ID
	if !model.CorrectPassword(password, int(userId)) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success."})
}
