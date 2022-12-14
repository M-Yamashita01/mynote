package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
)

type SignInParam struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func PostSignIn(c *gin.Context) {
	var signInParam SignInParam
	c.BindJSON(&signInParam)

	if existsUser(signInParam.Email) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Sign in failed."})
		return
	}

	user, err := model.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	userId := user.ID
	userProfile, err := model.CreateUserProfile(signInParam.FirstName, signInParam.LastName, signInParam.Email, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	if _, err := model.CreatePasswordAuthentication(signInParam.Password, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	userToken, err := model.CreateUserToken(userProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sign in failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": userToken.Token})
}

func existsUser(email string) bool {
	return model.ExistUserWithEmail(email)
}
