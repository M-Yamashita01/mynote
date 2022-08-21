package model_test

import (
	"MyNote/internal/model"
	"MyNote/pkg/database"
	"database/sql"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = BeforeSuite(func() {
	fmt.Println("beore suite")
})

var _ = Describe("PasswordAuthentication", Ordered, func() {
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

	Context("When password correct", func() {
		var userId uint

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId = user.Model.ID
			model.CreatePasswordAuthentication("password", userId)
		})

		It("return true.", func() {
			Expect(model.CorrectPassword("password", int(userId))).To(BeTrue())
		})
	})

	Context("When password incorrect", func() {
		var userId uint

		BeforeEach(func() {
			user, _ := model.CreateUser()
			userId = user.Model.ID
			model.CreatePasswordAuthentication("password", userId)
		})

		It("return false.", func() {
			Expect(model.CorrectPassword("IncorrectPassword", int(userId))).To(BeFalse())
		})
	})
})
