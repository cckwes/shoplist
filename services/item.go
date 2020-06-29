package services

import (
	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
)

func InsertItem(item *models.Item) error {
	err := db.DB.Create(item).Error
	return err
}

func FindItemByID(id string) (models.Item, error) {
	var item models.Item
	err := db.DB.Where("id = ?", id).First(&item).Error

	return item, err
}

func GetItemsByListID(listID string) ([]models.Item, error) {
	var items []models.Item
	err := db.DB.Where("list_id = ?", listID).Find(&items).Error

	return items, err
}
