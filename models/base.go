package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
	}
}
