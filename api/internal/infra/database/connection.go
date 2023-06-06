package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
