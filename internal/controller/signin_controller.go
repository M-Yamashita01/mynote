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

func SignIn(c *gin.Context) {
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if existsUser(email) {
		c.HTML(http.StatusBadRequest, "signin.html", gin.H{"err": "Already exists user with the email."})
		return
	}

	user, err := model.CreateUser()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signin.html", gin.H{"err": "Could not create account."})
		return
	}

	userId := user.Model.ID
	if _, err := model.CreateUserProfile(firstName, lastName, email, userId); err != nil {
		c.HTML(http.StatusInternalServerError, "signin.html", gin.H{"err": "Could not create account."})
		return
	}

	if _, err := model.CreatePasswordAuthentication(password, userId); err != nil {
		c.HTML(http.StatusInternalServerError, "signin.html", gin.H{"err": "Could not create account."})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func existsUser(email string) bool {
	return model.ExistUserWithEmail(email)
}
