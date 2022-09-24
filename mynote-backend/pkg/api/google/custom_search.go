package google

import (
	myNoteHttp "MyNote/pkg/http"
	myNoteOs "MyNote/pkg/os"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Article struct {
	Title    string
	SiteName string
	Url      string
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

func GetArticleSearchRequest(articleUrl string) (*Article, error) {
	apiKey := myNoteOs.GetEnv("API_KEY", "apiKey")
	cx := myNoteOs.GetEnv("CX", "cx")
	customSearchUrl := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", apiKey, cx, articleUrl)

	resp, err := myNoteHttp.GetRequest(new(http.Client), customSearchUrl)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	var customSearchResponseBody CustomSearchResponseBody
	json.Unmarshal(body, &customSearchResponseBody)

	article := createArticle(&customSearchResponseBody)
	return article, nil
}

func createArticle(customSearchResponseBody *CustomSearchResponseBody) *Article {
	article := Article{
		Title:    "",
		SiteName: "None",
		Url:      "",
	}

	item := customSearchResponseBody.Items[0]
	article.Title = item.Title
	for _, metatag := range item.PageMap.Metatags {
		if metatag.SiteName != "" {
			article.SiteName = metatag.SiteName
		}

		if metatag.Title != "" {
			article.Title = metatag.Title
		}
	}

	return &article
}
