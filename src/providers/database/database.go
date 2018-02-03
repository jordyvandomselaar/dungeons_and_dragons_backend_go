package database

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect(
	dialect string,
	username string,
	password string,
	database string,
) (*gorm.DB, error) {
	dbc, err := gorm.Open(dialect, username+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local")

	db = dbc

	return db, err
}

func Disconnect() {
	db.Close()
}

func GetDb() *gorm.DB {
	if db == nil {
		println("DB IS NIL.")
	}

	return db
}
