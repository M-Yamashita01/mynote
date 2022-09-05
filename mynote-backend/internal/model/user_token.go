package model

import (
	"MyNote/pkg/database"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	Token  string
	UserId int
}

func CreateUserToken(userProfile *UserProfile) (*UserToken, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &UserToken{}, err
	}

	defer database.Close(db)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        strconv.FormatUint(uint64(userProfile.UserId), 10),
		"firstName": userProfile.FirstName,
		"LastName":  userProfile.LastName,
		"email":     userProfile.Email,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("secret"))

	userToken := UserToken{
		Token:  tokenString,
		UserId: userProfile.UserId,
	}

	if err := db.Create(&userToken).Error; err != nil {
		log.Println("Error ocurred.")
		log.Println(err)
		return &UserToken{}, err
	}

	return &userToken, nil
}

func FindUserToken(userId int) (*UserToken, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &UserToken{}, err
	}

	defer database.Close(db)

	userToken := UserToken{}
	if err := db.Where("user_id = ?", userId).First(&userToken).Error; err != nil {
		return nil, err
	}

	return &userToken, nil
}

func FindUserId(token string) (int, error) {
	db, err := database.DbInit()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return -1, err
	}

	defer database.Close(db)

	userToken := UserToken{}
	if err := db.Where("token = ?", token).First(&userToken).Error; err != nil {
		return -1, err
	}

	return userToken.UserId, nil
}
