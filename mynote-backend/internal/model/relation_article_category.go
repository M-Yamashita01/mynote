package model

import "time"

type RelationArticleCategory struct {
	ID         uint `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	ArticleId  uint `gorm:"not null"`
	Article    Article
	CategoryId uint `gorm:"not null"`
	Category   Category
	CreatedAt  time.Time
}
