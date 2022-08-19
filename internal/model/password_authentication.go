package model

import (
	"log"

	"gorm.io/gorm"

	"MyNote/pkg/crypto"
	"MyNote/pkg/database"
)

type PasswordAuthentication struct {
	gorm.Model
	EncryptedPassword string
	UserId            int
}

func CreatePasswordAuthentication(password string, userId uint) (*PasswordAuthentication, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &PasswordAuthentication{}, err
	}

	defer database.Close(db)

	hashedPassword, _ := crypto.EncryptedPassword(password)
	passwordAuthentication := PasswordAuthentication{
		EncryptedPassword: hashedPassword,
		UserId:            int(userId),
	}

	if err := db.Create(&passwordAuthentication).Error; err != nil {
		log.Println("Error ocurred.")
		log.Println(err)
		return &PasswordAuthentication{}, err
	}

	return &passwordAuthentication, nil
}
