package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	myNoteOs "MyNote/pkg/os"
)

func DbInit() (*gorm.DB, error) {
	password := myNoteOs.GetEnv("MYSQL_ROOT_PASSWORD", "password")
	database := myNoteOs.GetEnv("MYSQL_DATABASE", "mynote")
	host := myNoteOs.GetEnv("MYSQL_HOST", "127.0.0.1")
	port := myNoteOs.GetEnv("MYSQL_PORT", "3306")

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", password, host, port, database)

	fmt.Println("dsn:" + dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
