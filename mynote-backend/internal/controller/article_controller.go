package controller

import (
	"MyNote/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetArticlesParam struct {
	SinceId      string `form:"since_id"`
	ArticleCount string `form:"article_count"`
}

type PostArticleParam struct {
	ArticleURL string `json:"article_url"`
}

func GetArticles(c *gin.Context) {
	var getArticlesParam GetArticlesParam
	if c.ShouldBindQuery(&getArticlesParam) != nil || getArticlesParam.SinceId == "" || getArticlesParam.ArticleCount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required params that are sinceId and articleCount"})
		return
	}

	articles, err := findArticleList(c.Request, getArticlesParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
	return
}

type Metatags struct {
	SiteName string `json:"site_name"`
	Title    string `json:"og:title"`
}

type PageMap struct {
	Metatags []Metatags `json:"metatags"`
}

type Items struct {
	Title   string  `json:"title"`
	PageMap PageMap `json:"pagemap"`
}

type CustomSearchResponseBody struct {
	Items []Items `json:"items"`
}

func PostArticle(c *gin.Context) {
	userId, err := model.FindUserIdFromRequestHeaderToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user"})
		return
	}

	var postArticleParam PostArticleParam
	c.BindJSON(&postArticleParam)

	if err := model.RegisterArticleFromUrl(postArticleParam.ArticleURL, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register article. err: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post article successfully."})
	return
}

func findArticleList(request *http.Request, getArticlesParam GetArticlesParam) ([](map[string]string), error) {
	userId, err := model.FindUserIdFromRequestHeaderToken(request)
	if err != nil {
		return nil, err
	}

	parsedSinceId, _ := strconv.ParseInt(getArticlesParam.SinceId, 10, 32)
	parsedArticleCount, _ := strconv.ParseInt(getArticlesParam.ArticleCount, 10, 32)

	articles, err := model.FindArticlesSinceId(userId, uint(parsedSinceId), int(parsedArticleCount))
	if err != nil {
		return nil, err
	}

	articleList := createArticleList(articles)

	return articleList, nil
}

func createArticleList(articles *[]model.Article) [](map[string]string) {
	articleList := make([](map[string]string), len((*articles)))

	for index, v := range *articles {
		articleMap := make(map[string]string)
		articleMap["article_id"] = strconv.FormatUint(uint64(v.ID), 10)
		articleMap["title"] = v.Title
		articleMap["url"] = v.Url
		articleMap["website_name"] = v.WebsiteName
		articleList[index] = articleMap
	}

	return articleList
}
