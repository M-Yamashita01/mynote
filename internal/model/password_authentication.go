package model

import (
	"gorm.io/gorm"
)

type PasswordAuthentication struct {
	gorm.Model
	encrypted_password string
	UserId             int
}
