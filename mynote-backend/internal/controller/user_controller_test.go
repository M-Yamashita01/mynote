package controller_test

import (
	"MyNote/internal/controller"
	"MyNote/internal/model"
	"MyNote/pkg/database/database_test"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", Ordered, func() {
	var db *database_test.DB

	BeforeAll(func() {
		db = database_test.ConnectTestDb()

		DeferCleanup(func() {
			db.CloseTestDb()
		})
	})

	Context("When get user request", func() {
		var userToken *model.UserToken
		var testContext *gin.Context
		var responseWriter *httptest.ResponseRecorder

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userProfile, _ := model.CreateUserProfile("first", "last", "example@example.com", user.ID)
			userToken, _ = model.CreateUserToken(userProfile)

			responseWriter = httptest.NewRecorder()
			testContext, _ = gin.CreateTestContext(responseWriter)
			testContext.Request, _ = http.NewRequest("GET", "/api/auth/user", nil)
			testContext.Request.Header.Set("Authorization", "Bearer "+userToken.Token)
		})

		It("Get user token successfully", func() {
			controller.GetUser(testContext)
			Expect(testContext.Writer.Status()).To(Equal(http.StatusOK))

			body, _ := ioutil.ReadAll(responseWriter.Body)

			type Token struct {
				Token string `json:"token"`
			}

			var responseUser struct {
				User Token `json:"user"`
			}

			json.Unmarshal([]byte(body), &responseUser)
			Expect(userToken.Token).To(Equal(responseUser.User.Token))
		})
	})
})
