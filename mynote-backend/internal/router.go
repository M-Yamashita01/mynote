package internal

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"MyNote/internal/controller"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.LoadHTMLGlob("web/*")
	r.Static("/web", "./web")

	r.POST("/api/auth/signin", controller.PostSignIn)
	r.POST("/api/auth/login", controller.PostLogin)
	r.GET("/api/auth/user", controller.GetUser)
	r.GET("/api/articles", controller.GetArticles)
	r.POST("/api/article", controller.PostArticle)

	return r
}
