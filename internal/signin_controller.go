package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MyNote/internal/model"
	"MyNote/pkg/crypto"
	"MyNote/pkg/database"
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

	db, _ := database.DbInit()
	user := model.User{}
	db.Create(&user)
	userId := user.Model.ID

	userProfile := model.UserProfile{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserId:    int(userId),
	}

	if db.Create(&userProfile).Error != nil {
		c.Abort()
		return
	}

	hashedPassword, _ := crypto.EncryptedPassword(password)
	passwordAuthentication := model.PasswordAuthentication{
		EncryptedPassword: hashedPassword,
		UserId:            int(userId),
	}

	if db.Create(&passwordAuthentication).Error != nil {
		c.Abort()
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
