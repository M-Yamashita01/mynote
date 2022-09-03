package internal

import (
	"github.com/gin-gonic/gin"

	"MyNote/internal/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("web/*")
	r.Static("/web", "./web")

	r.GET("/home", controller.ShowHome)
	r.GET("/signin", controller.ShowSignInPage())
	r.POST("/auth/api/signin", controller.PostSignIn)
	r.GET("/login", controller.ShowLoginPage())
	r.POST("/login", controller.PostLogin)

	return r
}
