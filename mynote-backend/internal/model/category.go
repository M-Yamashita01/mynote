package model

import "time"

type Category struct {
	ID                      uint   `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	CategoryName            string `gorm:"not null"`
	RelationArticleCategory []RelationArticleCategory
	UserId                  uint `gorm:"not null"`
	CreatedAt               time.Time
}
