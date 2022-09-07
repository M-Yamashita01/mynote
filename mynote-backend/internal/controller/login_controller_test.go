package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

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
			userId := user.ID
			userProfile, _ := model.CreateUserProfile("test_first_name", "test_last_name", "correct@example.com", userId)
			model.CreatePasswordAuthentication("CorrectPassword", userId)
			model.CreateUserToken(userProfile)
		})

		Context("Input correct email and password", func() {
			body := bytes.NewBufferString("{\"email\":\"correct@example.com\", \"password\":\"CorrectPassword\"}")

			responseWriter := httptest.NewRecorder()
			testContext, _ := gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/api/auth/login", body)
			testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)

			It("Login successfully", func() {
				controller.PostLogin(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		Context("Input incorrect email", func() {
			body := bytes.NewBufferString("{\"email\":\"incorrect@example.com\", \"password\":\"CorrectPassword\"}")

			responseWriter := httptest.NewRecorder()
			testContext, _ := gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/api/auth/login", body)
			testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)

			It("Login failed", func() {
				controller.PostLogin(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusUnauthorized))
			})
		})

		Context("Input incorrect password", func() {
			body := bytes.NewBufferString("{\"email\":\"correct@example.com\", \"password\":\"IncorrectPassword\"}")

			responseWriter := httptest.NewRecorder()
			testContext, _ := gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/api/auth/login", body)
			testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)

			It("Login failed", func() {
				controller.PostLogin(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
