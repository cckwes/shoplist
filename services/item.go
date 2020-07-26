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

func UpdateItem(item models.Item) (models.Item, error) {
	if err := db.DB.Model(item).Where("id = ?", item.ID).Updates(map[string]interface{}{"name": item.Name, "count": item.Count, "done": item.Done, "removed": item.Removed}).Error; err != nil {
		return models.Item{}, err
	}

	return FindItemByID(item.ID)
}
