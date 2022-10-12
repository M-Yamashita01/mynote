package migrate

import (
	"log"

	"MyNote/internal/model"
	"MyNote/pkg/database"
)

func SetupDb() {
	db, err := database.ConnectDb()
	if err != nil {
		log.Println(err)
		return
	}

	db.AutoMigrate(&(model.User{}))
	db.AutoMigrate(&(model.UserProfile{}))
	db.AutoMigrate(&(model.PasswordAuthentication{}))
	db.AutoMigrate(&(model.UserToken{}))
	db.AutoMigrate(&(model.Article{}))
	db.AutoMigrate(&(model.Category{}))
	db.AutoMigrate(&(model.RelationArticleCategory{}))
}
