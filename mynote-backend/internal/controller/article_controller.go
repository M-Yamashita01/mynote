package controller

import (
	"MyNote/internal/model"
	"encoding/json"
	"fmt"
	"io"
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

type CustomSearchResult struct {
	Title string `form:"title"`
}

type CustomSearchResponseBody struct {
	Items []CustomSearchResult `form:"items"`
}

func PostArticle(c *gin.Context) {
	userId, err := model.FindUserIdFromRequestHeaderToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user"})
		return
	}

	var postArticleParam PostArticleParam
	c.BindJSON(&postArticleParam)

	apiKey := ""
	cx := ""
	customSearchUrl := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", apiKey, cx, postArticleParam.ArticleURL)

	req, err := http.NewRequest(http.MethodGet, customSearchUrl, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid URL. err: " + err.Error()})
		return
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Request. err: " + err.Error()})
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Response. StatusCode: " + fmt.Sprint(resp.StatusCode) + "\nerr: " + err.Error()})
		return
	}

	body, _ := io.ReadAll(resp.Body)
	var customSearchResponseBody CustomSearchResponseBody
	json.Unmarshal(body, &customSearchResponseBody)

	title := customSearchResponseBody.Items[0].Title
	model.CreateArticle(title, postArticleParam.ArticleURL, "", userId)

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
