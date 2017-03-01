package handler

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type APIS struct {
	ID     uint `gorm:"primary_key"`
	Method string
	Path   string
	Value  string
}

var gDb *gorm.DB

func Connect(path string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	gDb = db
	return db, err
}

func GetDB() *gorm.DB {
	return gDb
}
