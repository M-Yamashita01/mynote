package controller

import (
	"MyNote/internal/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetArticlesParam struct {
	SinceId      string `form:"since_id"`
	ArticleCount string `form:"article_count"`
}

func GetArticles(c *gin.Context) {
	header := c.Request.Header
	bearToken := header["Authorization"]
	splitBearToken := strings.Split(bearToken[0], " ")
	token := splitBearToken[1]

	var getArticlesParam GetArticlesParam
	if c.ShouldBindQuery(&getArticlesParam) != nil || getArticlesParam.SinceId == "" || getArticlesParam.ArticleCount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required params that are sinceId and articleCount"})
		return
	}

	userId, err := model.FindUserId(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user"})
		return
	}

	parsedSinceId, _ := strconv.ParseInt(getArticlesParam.SinceId, 10, 32)
	parsedArticleCount, _ := strconv.ParseInt(getArticlesParam.ArticleCount, 10, 32)

	articles, err := model.FindArticlesSinceId(userId, uint(parsedSinceId), int(parsedArticleCount))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not found articles"})
		return
	}

	articleList := make([](map[string]string), len((*articles)))
	for index, v := range *articles {
		articleMap := make(map[string]string)
		articleMap["article_id"] = strconv.FormatUint(uint64(v.ID), 10)
		articleMap["title"] = v.Title
		articleMap["url"] = v.Url
		articleMap["website_name"] = v.WebsiteName
		articleList[index] = articleMap
	}

	c.JSON(http.StatusOK, gin.H{"articles": articleList})
	return
}
