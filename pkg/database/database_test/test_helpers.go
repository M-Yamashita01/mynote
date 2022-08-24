package database_test

import (
	"os"

	"gorm.io/gorm"

	"MyNote/internal/model"
	"MyNote/pkg/database"
)

type DB struct {
	testDb    string
	currentDb string
	GormDb    *gorm.DB
}

func ConnectTestDb() *DB {
	var db DB

	db.testDb = "mynote_test_db"
	db.currentDb = "mynote_db"
	os.Setenv("MYSQL_ROOT_PASSWORD", "password")
	os.Setenv("MYSQL_DATABASE", db.testDb)

	db.GormDb, _ = database.DbInit()

	db.migrate()

	return &db
}

func (db *DB) migrate() {
	db.GormDb.AutoMigrate(&(model.User{}))
	db.GormDb.AutoMigrate(&(model.UserProfile{}))
	db.GormDb.AutoMigrate(&(model.PasswordAuthentication{}))
}

func (db *DB) CloseTestDb() {
	db.GormDb.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.GormDb.Exec("TRUNCATE user_profiles;")
	db.GormDb.Exec("TRUNCATE password_authentications;")
	db.GormDb.Exec("TRUNCATE users;")
	db.GormDb.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	sqlDb, _ := db.GormDb.DB()
	sqlDb.Close()

	os.Setenv("MYSQL_DATABASE", db.currentDb)
}
