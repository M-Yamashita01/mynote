package http_test

import (
	"MyNote/pkg/database/database_test"
	myNoteHttp "MyNote/pkg/http"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Describe("GetBearerTokenFromHeader", Ordered, func() {
		responseWriter := httptest.NewRecorder()
		testContext, _ := gin.CreateTestContext(responseWriter)
		testContext.Request, _ = http.NewRequest("GET", "http://localhost:8080/api/articles?since_id=0&article_count=10", nil)
		testContext.Request.Header.Set("Authorization", "Bearer tokenexample")

		It("Get token successfully", func() {
			token := myNoteHttp.GetBearerTokenFromHeader(testContext.Request)
			Expect(token).To(Equal("tokenexample"))
		})
	})
})
