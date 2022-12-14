package model

import (
	"log"
	"time"

	"MyNote/pkg/database"
)

type User struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Article   Article
	Category  Category
	CreatedAt time.Time
}

func CreateUser() (*User, error) {
	db, err := database.ConnectDb()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &User{}, err
	}

	defer database.Close(db)

	var user User
	if err := db.Create(&user).Error; err != nil {
		log.Println("Error ocurred.")
		log.Println(err)
		return &User{}, err
	}

	return &user, nil
}
