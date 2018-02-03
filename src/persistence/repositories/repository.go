package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	db *gorm.DB
}
