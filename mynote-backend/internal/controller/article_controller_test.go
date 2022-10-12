package controller_test

import (
	"MyNote/internal/controller"
	"MyNote/internal/model"
	"MyNote/pkg/api/google"
	"MyNote/pkg/database/database_test"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArticleController", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Describe("GetArticles", Ordered, func() {
		Context("When two article exist", func() {
			var userId uint
			var testContext *gin.Context
			var responseWriter *httptest.ResponseRecorder

			BeforeEach(func() {
				user, _ := model.CreateUser()
				userId = user.ID
				model.CreateArticle("TestTitle", "https://www.example.com/", "example.com", userId)
				model.CreateArticle("TestTitle2", "https://www.example.com/", "example.com", userId)

				userProfile, _ := model.CreateUserProfile("test", "last", "example@example.com", userId)
				userToken, _ := model.CreateUserToken(userProfile)

				responseWriter = httptest.NewRecorder()
				testContext, _ = gin.CreateTestContext(responseWriter)
				testContext.Request, _ = http.NewRequest("GET", "http://localhost:8080/api/articles?since_id=0&article_count=10", nil)
				testContext.Request.Header.Set("Authorization", "Bearer "+userToken.Token)
			})

			It("Get articles successfully", func() {
				controller.GetArticles(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusOK))

				body, _ := ioutil.ReadAll(responseWriter.Body) // response body is []byte

				type Article struct {
					ArticleId   string `json:"article_id"`
					Title       string `json:"title"`
					Url         string `json:"url"`
					WebsiteName string `json:"website_name"`
				}

				var articles struct {
					Articles []Article `json:"articles"`
				}

				json.Unmarshal([]byte(body), &articles)
				arrayArticles := articles.Articles

				firstArticle := arrayArticles[0]
				Expect(firstArticle.Title).To(Equal("TestTitle2"))
				Expect(firstArticle.Url).To(Equal("https://www.example.com/"))
				Expect(firstArticle.WebsiteName).To(Equal("example.com"))

				secondArticle := arrayArticles[1]
				Expect(secondArticle.Title).To(Equal("TestTitle"))
				Expect(secondArticle.Url).To(Equal("https://www.example.com/"))
				Expect(secondArticle.WebsiteName).To(Equal("example.com"))
			})
		})

		Context("When send request with no parameters", func() {
			var userId uint
			var testContext *gin.Context
			var responseWriter *httptest.ResponseRecorder

			BeforeEach(func() {
				user, _ := model.CreateUser()
				userId = user.ID
				userProfile, _ := model.CreateUserProfile("test", "last", "example@example.com", userId)
				userToken, _ := model.CreateUserToken(userProfile)

				responseWriter = httptest.NewRecorder()
				testContext, _ = gin.CreateTestContext(responseWriter)
				testContext.Request, _ = http.NewRequest("GET", "http://localhost:8080/api/articles", nil)
				testContext.Request.Header.Set("Authorization", "Bearer "+userToken.Token)
			})

			It("Get BadRequest status", func() {
				controller.GetArticles(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			})
		})

		Context("When send request with wrong parameters", func() {
			var userId uint
			var testContext *gin.Context
			var responseWriter *httptest.ResponseRecorder

			BeforeEach(func() {
				user, _ := model.CreateUser()
				userId = user.ID
				userProfile, _ := model.CreateUserProfile("test", "last", "example@example.com", userId)
				userToken, _ := model.CreateUserToken(userProfile)

				responseWriter = httptest.NewRecorder()
				testContext, _ = gin.CreateTestContext(responseWriter)
				testContext.Request, _ = http.NewRequest("GET", "http://localhost:8080/api/articles?test_id=0&test_count=10", nil)
				testContext.Request.Header.Set("Authorization", "Bearer "+userToken.Token)
			})

			It("Get BadRequest status", func() {
				controller.GetArticles(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			})
		})

	})

	Describe("PostArticle", func() {
		Context("When post correct article's url", func() {
			var userId uint
			var testContext *gin.Context
			var responseWriter *httptest.ResponseRecorder
			var strCustomSearchResponseBody string

			BeforeEach(func() {
				body := bytes.NewBufferString(
					"{\"article_url\":\"https://example.com/\"}",
				)

				user, _ := model.CreateUser()
				userId = user.ID
				userProfile, _ := model.CreateUserProfile("test", "last", "example@example.com", userId)
				userToken, _ := model.CreateUserToken(userProfile)

				responseWriter = httptest.NewRecorder()
				testContext, _ = gin.CreateTestContext(responseWriter)
				testContext.Request, _ = http.NewRequest("POST", "http://localhost:8080/api/article", body)
				testContext.Request.Header.Set("Authorization", "Bearer "+userToken.Token)
				testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)

				customSearchResponseBody := &google.CustomSearchResponseBody{
					Items: []google.Items{
						{
							Title: "item title",
							PageMap: google.PageMap{
								Metatags: []google.Metatags{
									{
										Title:    "metatag title",
										SiteName: "site name",
									},
								},
							},
						},
					},
				}
				jsonCustomSearchResponseBody, _ := json.Marshal(&customSearchResponseBody)
				strCustomSearchResponseBody = string(jsonCustomSearchResponseBody)
			})

			It("Get articles successfully", func() {
				httpmock.Activate()
				defer httpmock.DeactivateAndReset()

				r := regexp.MustCompile(`(?m)https://www.googleapis.com/customsearch/v1([a-zA-Z0-9:\/_\=\-\+\&\?\.]{1,})$`)
				httpmock.RegisterRegexpResponder("GET", r,
					httpmock.NewStringResponder(http.StatusOK, strCustomSearchResponseBody))

				controller.PostArticle(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})
	})
})
