package internal

import (
	"github.com/gin-gonic/gin"

	"MyNote/internal/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/*.html")

	// r.GET("/", ShowList())
	r.GET("/signin", controller.ShowSignInPage())
	r.POST("/signin", controller.SignIn)
	r.GET("/login", controller.ShowLoginPage())
	r.POST("/login", controller.PostLogin)

	return r
}
