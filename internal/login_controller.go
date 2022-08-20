package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
	"MyNote/model"
)

func ShowLoginPage() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	}
}

func PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	// password := c.PostForm("psw")

	if !model.ExistUserWithEmail(email) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": "Email or password is incorrect."})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
