package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"MyNote/internal/controller"
	"MyNote/internal/model"
	"MyNote/pkg/database/database_test"
)

// "MyNote/pkg/database"

var _ = Describe("LoginController", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	It("Show LogInPage successfully", func() {
		controller.ShowLoginPage()
	})

	Describe("PostLogIn", Ordered, func() {
		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId := user.Model.ID
			model.CreateUserProfile("test_first_name", "test_last_name", "example@example.com", userId)
			model.CreatePasswordAuthentication("ExamplePassword", userId)
		})

		Context("Input correct email and password", func() {
			form := url.Values{}
			form.Add("email", "example@example.com")
			form.Add("password", "ExamplePassword")
			body := strings.NewReader(form.Encode())

			responseWriter := httptest.NewRecorder()
			testContext, _ := gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/login", body)
			testContext.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			It("Login successfully", func() {
				controller.PostLogin(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusFound))
			})
		})
	})
})
