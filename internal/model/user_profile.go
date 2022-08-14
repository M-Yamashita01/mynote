package model

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	firstName string
	lastName  string
	email     string
	UserId    int
}
