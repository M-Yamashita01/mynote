package controller

import (
	"MyNote/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := model.FindUserIdFromRequestHeaderToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userId})
}
