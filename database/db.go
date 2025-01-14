package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB


 type Album struct{
	gorm.Model
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}




func InitDB() error {

	var err error

	db,err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
			return err
	}

		return nil
}

func GetDb() *gorm.DB {
	return db
}