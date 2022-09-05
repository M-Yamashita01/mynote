package controller

import (
	"MyNote/internal/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	header := c.Request.Header
	bearToken := header["Authorization"]
	splitedBearToken := strings.Split(bearToken[0], " ")
	token := splitedBearToken[1]
	userId, err := model.FindUserId(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userId})
}
