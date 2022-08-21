package internal_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"MyNote/internal"
	"MyNote/internal/model"
	"MyNote/pkg/database"
)

// "MyNote/pkg/database"

var _ = Describe("LoginController", Ordered, func() {
	var currentdb string
	var db *gorm.DB
	var sqlDb *sql.DB

	BeforeAll(func() {
		testdb := "mynote_test_db"
		currentdb = "mynote_db"
		os.Setenv("MYSQL_ROOT_PASSWORD", "password")
		os.Setenv("MYSQL_DATABASE", testdb)
		db, _ = database.DbInit()
		sqlDb, _ = db.DB()

		DeferCleanup(func() {
			db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
			db.Exec("TRUNCATE user_profiles;")
			db.Exec("TRUNCATE password_authentications;")
			db.Exec("TRUNCATE users;")
			db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
			sqlDb.Close()
			os.Setenv("MYSQL_DATABASE", currentdb)
		})
	})

	It("Show LogInPage successfully", func() {
		internal.ShowLoginPage()
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
				internal.PostLogin(testContext)
				Expect(testContext.Writer.Status()).To(Equal(http.StatusFound))
			})
		})
	})
})
