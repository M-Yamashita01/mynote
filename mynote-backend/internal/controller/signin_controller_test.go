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

var _ = Describe("SignInController", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Context("When posts correct signin request", func() {
		var testContext *gin.Context

		BeforeEach(func() {
			body := bytes.NewBufferString(
				"{\"firstName\":\"test_first_name\",\"lastName\":\"test_last_name\",\"email\": \"example@example.com\", \"password\": \"ExamplePassword\"}",
			)

			responseWriter := httptest.NewRecorder()
			testContext, _ = gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/api/auth/signin", body)
			testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)
		})

		AfterEach(func() {
			db.TruncateAllTable()
		})

		It("Get 200 status code.", func() {
			controller.PostSignIn(testContext)

			Expect(testContext.Writer.Status()).To(Equal(http.StatusOK))
		})

		It("Create user profile record.", func() {
			controller.PostSignIn(testContext)

			var userProfile model.UserProfile
			db.GormDb.Last(&userProfile)

			Expect(userProfile.FirstName).To(Equal("test_first_name"))
			Expect(userProfile.LastName).To(Equal("test_last_name"))
			Expect(userProfile.Email).To(Equal("example@example.com"))
		})

		It("Create password authentication record.", func() {
			controller.PostSignIn(testContext)

			var passwordAuthentication model.PasswordAuthentication
			db.GormDb.Last(&passwordAuthentication)

			Expect(passwordAuthentication.ID).NotTo(BeNil())
			Expect(passwordAuthentication.EncryptedPassword).NotTo(BeNil())
		})
	})

	Context("When user with same password exists", func() {
		var testContext *gin.Context

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId := user.ID
			model.CreateUserProfile("test_first_name", "test_last_name", "example@example.com", userId)

			body := bytes.NewBufferString(
				"{\"firstName\":\"test_first_name\",\"lastName\":\"test_last_name\",\"email\": \"example@example.com\", \"password\": \"ExamplePassword\"}",
			)

			responseWriter := httptest.NewRecorder()
			testContext, _ = gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("POST", "/api/auth/signin", body)
			testContext.Request.Header.Set("Content-Type", gin.MIMEJSON)
		})

		AfterEach(func() {
			db.TruncateAllTable()
		})

		It("Get 401 status code.", func() {
			controller.PostSignIn(testContext)

			Expect(testContext.Writer.Status()).To(Equal(http.StatusUnauthorized))
		})

		It("Not create user profile record.", func() {
			controller.PostSignIn(testContext)

			var userProfile model.UserProfile
			result := db.GormDb.Find(&userProfile)

			Expect(int(result.RowsAffected)).To(Equal(1))
		})
	})
})
