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

var _ = Describe("SignInController", func() {
	var currentdb string
	var db *gorm.DB
	var sqlDb *sql.DB

	BeforeEach(func() {
		testdb := "mynote_test_db"
		currentdb = "mynote_db"
		os.Setenv("MYSQL_ROOT_PASSWORD", "password")
		os.Setenv("MYSQL_DATABASE", testdb)
		db, _ = database.DbInit()
		sqlDb, _ = db.DB()

		db.AutoMigrate(&(model.User{}))
		db.AutoMigrate(&(model.UserProfile{}))
		db.AutoMigrate(&(model.PasswordAuthentication{}))
	})

	AfterEach(func() {
		sqlDb.Exec("SET FOREIGN_KEY_CHECKS = 0;")
		sqlDb.Exec("TRUNCATE user_profiles;")
		sqlDb.Exec("TRUNCATE password_authentications;")
		sqlDb.Exec("TRUNCATE users;")
		sqlDb.Exec("SET FOREIGN_KEY_CHECKS = 1;")
		sqlDb.Close()
		os.Setenv("MYSQL_DATABASE", currentdb)
	})

	Context("When posts signin request", func() {
		form := url.Values{}
		form.Add("firstName", "test_first_name")
		form.Add("lastName", "test_last_name")
		form.Add("email", "example@example.com")
		form.Add("password", "ExamplePassword")
		body := strings.NewReader(form.Encode())

		testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
		testContext.Request, _ = http.NewRequest("POST", "/signin", body)
		testContext.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		It("Create user profile successfully", func() {
			internal.SignIn(testContext)

			var userProfile model.UserProfile
			db.First(&userProfile)

			Expect(userProfile.FirstName).To(Equal("test_first_name"))
			Expect(userProfile.LastName).To(Equal("test_last_name"))
			Expect(userProfile.Email).To(Equal("example@example.com"))
		})

		It("Create password authentication successfully", func() {
			internal.SignIn(testContext)

			var passwordAuthentication model.PasswordAuthentication
			db.First(&passwordAuthentication)

			Expect(passwordAuthentication.Model.ID).NotTo(BeNil())
			Expect(passwordAuthentication.EncryptedPassword).NotTo(BeNil())
		})
	})
})
