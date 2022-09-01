package model

import (
	"log"

	"gorm.io/gorm"

	"MyNote/pkg/database"
)

type UserProfile struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	UserId    int
}

func CreateUserProfile(firstName string, lastName string, email string, userId uint) (*UserProfile, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &UserProfile{}, err
	}

	defer database.Close(db)

	userProfile := UserProfile{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserId:    int(userId),
	}

	if err := db.Create(&userProfile).Error; err != nil {
		log.Println("Error ocurred.")
		log.Println(err)
		return &UserProfile{}, err
	}

	return &userProfile, nil
}

func ExistUserWithEmail(email string) bool {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return false
	}

	if (int(db.Where("email = ?", email).Take(&UserProfile{}).RowsAffected) == 0) {
		return false
	}

	return true
}

func FindUserProfile(email string) (*UserProfile, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &UserProfile{}, err
	}

	defer database.Close(db)

	userProfile := UserProfile{}
	if err := db.Where("email = ?", email).First(&userProfile).Error; err != nil {
		return nil, err
	}

	return &userProfile, nil
}
