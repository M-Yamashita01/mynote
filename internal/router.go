package internal

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/*.html")

	// r.GET("/", ShowList())
	r.GET("/signin", ShowSignInPage())
	r.POST("/signin", SignIn)
	r.GET("/login", ShowLoginPage())
	r.GET("/login", PostLogin)

	return r
}
