package services

import (
	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
)

func FindListsByUserID(userID string) []models.List {
	var lists []models.List

	db.DB.Preload("Items").Find(&lists).Where("user_id = ?", userID)

	return lists
}

func GetListByID(id string) (models.List, error) {
	var list models.List
	err := db.DB.Where("id = ?", id).First(&list).Error

	return list, err
}

func InsertList(list *models.List) error {
	err := db.DB.Create(list).Error
	return err
}

func UpdateList(list models.List) (models.List, error) {
	if err := db.DB.Model(list).Where("id = ?", list.ID).Update("name", list.Name).Error; err != nil {
		return models.List{}, err
	}
	return GetListByID(list.ID)
}