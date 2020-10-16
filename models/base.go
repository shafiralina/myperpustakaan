package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/perpustakaan")
	if err != nil {
		log.Println(err)
	}
	return db
}
