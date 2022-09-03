package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
)

func ShowSignInPage() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "signin.html", nil)
	}
}

func PostSignIn(c *gin.Context) {
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if existsUser(email) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Sign in failed."})
		return
	}

	user, err := model.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	userId := user.Model.ID
	if _, err := model.CreateUserProfile(firstName, lastName, email, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	if _, err := model.CreatePasswordAuthentication(password, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign in success."})
}

func existsUser(email string) bool {
	return model.ExistUserWithEmail(email)
}
