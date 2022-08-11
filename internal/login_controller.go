package internal

import "github.com/gin-gonic/gin"

func ShowLoginPage() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	}
}
