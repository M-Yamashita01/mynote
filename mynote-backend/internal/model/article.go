package model

import "time"

type Article struct {
	ID                      uint `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Title                   string
	Url                     string
	WebsiteName             string
	RelationArticleCategory []RelationArticleCategory
	UserId                  uint `gorm:"not null"`
	CreatedAt               time.Time
}
