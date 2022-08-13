package migrate

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"MyNote/internal/model"
)

func SetupDb() {
	db, err := dbInit()
	if err != nil {
		return
	}

	db.AutoMigrate(&(model.User{}))
	db.AutoMigrate(&(model.UserProfile{}))
	db.AutoMigrate(&(model.PasswordAuthentication{}))
}

func dbInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/mynote_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
