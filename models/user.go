package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Base
	Name     string
	Email    string
	ShopList []Item
}
