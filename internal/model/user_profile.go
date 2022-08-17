package model

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	UserId    int
}
