package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/mynote_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
