package model

import (
	"MyNote/pkg/database"
	myNoteHttp "MyNote/pkg/http"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserToken struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Token     string `gorm:"not null"`
	UserId    uint   `gorm:"not null"`
	User      User
	CreatedAt time.Time
}

func CreateUserToken(userProfile *UserProfile) (*UserToken, error) {
	db, err := database.ConnectDb()
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

func FindUserToken(userId uint) (*UserToken, error) {
	db, err := database.ConnectDb()
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

func FindUserIdFromRequestHeaderToken(request *http.Request) (uint, error) {
	token := myNoteHttp.GetBearerTokenFromHeader(request)

	db, err := database.ConnectDb()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return 0, err
	}

	defer database.Close(db)

	userToken := UserToken{}
	if err := db.Where("token = ?", token).First(&userToken).Error; err != nil {
		return 0, err
	}

	return userToken.UserId, nil
}
