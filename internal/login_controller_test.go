package internal_test

import (
	"database/sql"
	"os"

	. "github.com/onsi/ginkgo/v2"
	"gorm.io/gorm"

	"MyNote/internal"
	"MyNote/pkg/database"
)

// "MyNote/pkg/database"

var _ = Describe("LoginController", func() {
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
	})

	AfterEach(func() {
		sqlDb.Exec("truncate users;")
		sqlDb.Close()
		os.Setenv("MYSQL_DATABASE", currentdb)
	})

	It("Show LogInPage successfully", func() {
		internal.ShowLoginPage()
	})
})
