package services

import (
	db "github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
)

func GetOrCreateUser(email string) (models.User, error) {
	user := models.User{}
	err := db.DB.FirstOrCreate(&user, models.User{Email: email}).Error

	return user, err
}
