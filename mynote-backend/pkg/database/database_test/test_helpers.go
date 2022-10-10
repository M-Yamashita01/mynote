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

	db.GormDb, _ = database.ConnectDb()

	db.migrate()

	return &db
}

func (db *DB) migrate() {
	db.GormDb.AutoMigrate(&(model.User{}))
	db.GormDb.AutoMigrate(&(model.UserProfile{}))
	db.GormDb.AutoMigrate(&(model.PasswordAuthentication{}))
	db.GormDb.AutoMigrate(&(model.UserToken{}))
	db.GormDb.AutoMigrate(&(model.Article{}))
	db.GormDb.AutoMigrate(&(model.Category{}))
	db.GormDb.AutoMigrate(&(model.RelationArticleCategory{}))
}

func (db *DB) CloseTestDb() {
	db.TruncateAllTable()

	sqlDb, _ := db.GormDb.DB()
	sqlDb.Close()

	os.Setenv("MYSQL_DATABASE", db.currentDb)
}

func (db *DB) TruncateAllTable() {
	db.GormDb.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.GormDb.Exec("TRUNCATE relation_article_categories;")
	db.GormDb.Exec("TRUNCATE categories;")
	db.GormDb.Exec("TRUNCATE articles;")
	db.GormDb.Exec("TRUNCATE user_tokens;")
	db.GormDb.Exec("TRUNCATE user_profiles;")
	db.GormDb.Exec("TRUNCATE password_authentications;")
	db.GormDb.Exec("TRUNCATE users;")
	db.GormDb.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}
