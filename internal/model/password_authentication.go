package model

import (
	"gorm.io/gorm"
)

type PasswordAuthentication struct {
	gorm.Model
	EncryptedPassword string
	UserId            int
}
