package model_test

import (
	"MyNote/internal/model"
	"MyNote/pkg/database/database_test"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Article", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Describe("FindLatestArticles", Ordered, func() {
		Context("When an article exist", func() {
			var userId uint

			BeforeEach(func() {
				user, _ := model.CreateUser()
				userId = user.ID
				model.CreateArticle("TestTitle", "https://www.example.com/", "example.com", userId)
				model.CreateArticle("TestTitle2", "https://www.example.com/", "example.com", userId)
			})

			AfterEach(func() {
				db.TruncateAllTable()
			})

			It("return two records of articles.", func() {
				sinceId := 0
				articleCount := 10
				articles, _ := model.FindArticlesSinceId(userId, uint(sinceId), articleCount)
				Expect(len((*articles))).To(Equal(2))

				firstArticle := (*articles)[0]
				Expect(firstArticle.Title).To(Equal("TestTitle2"))
				secondArticle := (*articles)[1]
				Expect(secondArticle.Title).To(Equal("TestTitle"))
			})

			Context("specified sinceId greater than zero", func() {
				It("return articles with ids greater than sinceId.", func() {
					sinceId := 1
					articleCount := 10
					articles, _ := model.FindArticlesSinceId(userId, uint(sinceId), articleCount)
					Expect(len((*articles))).To(Equal(1))

					article := (*articles)[0]
					Expect(article.Title).To(Equal("TestTitle2"))
				})
			})
			// TODO Write tests for checking limit 10 over records
		})
	})
})
