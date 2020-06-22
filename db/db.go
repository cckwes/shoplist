package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Open() error {
	var err error
	DB, err = gorm.Open("sqlite3", "data.db")
	return err
}

func Close() error {
	return DB.Close()
}
