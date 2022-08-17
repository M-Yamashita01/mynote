package internal

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

	user, err := model.CreateUser()
	if err != nil {
		c.Abort()
		return
	}

	userId := user.Model.ID
	if _, err := model.CreateUserProfile(firstName, lastName, email, userId); err != nil {
		c.Abort()
		return
	}

	if _, err := model.CreatePasswordAuthentication(password, userId); err != nil {
		c.Abort()
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
