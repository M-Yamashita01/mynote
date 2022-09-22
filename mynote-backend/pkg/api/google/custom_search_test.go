package google_test

import (
	"MyNote/pkg/api/google"
	"MyNote/pkg/database/database_test"

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
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://www.googleapis.com/customsearch/v1",
			httpmock.NewStringResponder(200, "mocked"),
		)

		It("test", func() {
			_, err := google.GetArticleSearchRequest("")
			Expect(err).To(BeNil())
		})
	})
})
