package google_test

import (
	"MyNote/pkg/api/google"
	"MyNote/pkg/database/database_test"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomSearch", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Describe("GetArticleSearchRequest", Ordered, func() {
		Context("article has metatags", func() {
			body := &google.CustomSearchResponseBody{
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
			jsonBody, _ := json.Marshal(&body)

			It("Get article title and sitename from ogp.", func() {
				httpmock.Activate()
				defer httpmock.DeactivateAndReset()

				r := regexp.MustCompile(`(?m)https://www.googleapis.com/customsearch/v1([a-zA-Z0-9:\/_\=\-\+\&\?\.]{1,})$`)
				httpmock.RegisterRegexpResponder("GET", r,
					httpmock.NewStringResponder(http.StatusOK, string(jsonBody)))

				article, err := google.GetArticleSearchRequest("https://example.com")
				Expect(err).To(BeNil())
				Expect(article.Title).To(Equal("metatag title"))
				Expect(article.SiteName).To(Equal("site name"))
			})
		})

		Context("article does not have metatag, but has items", func() {
			body := &google.CustomSearchResponseBody{
				Items: []google.Items{
					{
						Title:   "item title",
						PageMap: google.PageMap{},
					},
				},
			}
			jsonBody, _ := json.Marshal(&body)

			It("Get article title from ogp and default site name", func() {
				httpmock.Activate()
				defer httpmock.DeactivateAndReset()

				r := regexp.MustCompile(`(?m)https://www.googleapis.com/customsearch/v1([a-zA-Z0-9:\/_\=\-\+\&\?\.]{1,})$`)
				httpmock.RegisterRegexpResponder("GET", r,
					httpmock.NewStringResponder(http.StatusOK, string(jsonBody)))

				article, err := google.GetArticleSearchRequest("https://example.com")
				Expect(err).To(BeNil())
				Expect(article.Title).To(Equal("item title"))
				Expect(article.SiteName).To(Equal("None"))
			})
		})

	})
})
