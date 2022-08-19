package model

import (
	"log"

	"gorm.io/gorm"

	"MyNote/pkg/database"
)

type User struct {
	gorm.Model
	UserProfile            UserProfile
	PasswordAuthentication PasswordAuthentication
}

func CreateUser() (*User, error) {
	db, err := database.DbInit()
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
