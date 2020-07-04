package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Base
	Email string `json:"email"`
	Lists []List `json:"lists"`
}
