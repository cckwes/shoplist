package services

import (
	db "github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
)

func CreateUser(email string) {
	var user models.User
	user.Email = email
	db.DB.Create(user)
}
