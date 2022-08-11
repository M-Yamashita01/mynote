package internal

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/*.html")

	r.GET("/", ShowList())
	r.GET("/login", ShowLoginPage())

	return r
}
