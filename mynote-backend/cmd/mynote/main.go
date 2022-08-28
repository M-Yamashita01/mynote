package main

import (
	"MyNote/internal"
	"MyNote/internal/migrate"
)

func main() {
	migrate.SetupDb()

	router := internal.GetRouter()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// "message": "pong",
	// })
	// })
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
