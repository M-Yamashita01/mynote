package model

import (
	"MyNote/pkg/database"
	"log"
	"time"
)

type Article struct {
	ID                      uint `gorm:"primary_key;AUTO_INCREMENT;not null;"`
	Title                   string
	Url                     string
	WebsiteName             string
	RelationArticleCategory []RelationArticleCategory
	UserId                  uint `gorm:"not null"`
	CreatedAt               time.Time
}

func CreateArticle(title string, url string, webSiteName string, userId uint) (*Article, error) {
	db, err := database.ConnectDb()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return &Article{}, err
	}

	defer database.Close(db)

	article := Article{
		Title:       title,
		Url:         url,
		WebsiteName: webSiteName,
		UserId:      userId,
	}

	if err := db.Create(&article).Error; err != nil {
		log.Println("Error ocurred.")
		log.Println(err)
		return &Article{}, err
	}

	return &article, nil
}

func FindArticlesSinceId(userId uint, sinceId uint, articleCount int) (*[]Article, error) {
	db, err := database.ConnectDb()
	if err != nil {
		log.Println("Failed db connection.")
		log.Println(err)
		return nil, err
	}

	defer database.Close(db)

	articles := []Article{}
	if err := db.Where("user_id = ? AND id > ?", userId, sinceId).Order("id desc").Limit(articleCount).Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}
