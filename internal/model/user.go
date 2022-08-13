package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserProfile            UserProfile
	PasswordAuthentication PasswordAuthentication
}
