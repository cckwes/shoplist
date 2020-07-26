package controllers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/services"
)

type CreateItemInput struct {
	Name   string `json:"name"`
	Count  uint   `json:"count"`
	ListID string `json:"list_id"`
}

type UpdateItemInput struct {
	Name    string `json:"name"`
	Count   uint   `json:"count"`
	Done    bool   `json:"done"`
	Removed bool   `json:"removed"`
}

func CreateItem(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID

	var input CreateItemInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input", err)
		c.Status(400)
		return
	}

	name := strings.TrimSpace(input.Name)
	if len(name) == 0 {
		log.Println("Received invalid name for create item")
		c.Status(400)
		return
	}
	if input.Count <= 0 {
		log.Println("Receive count less than or equal to zero")
		c.Status(400)
		return
	}

	list, err := services.GetListByID(input.ListID)
	if err != nil {
		log.Println("Failed to find list", err)
		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if list.UserID != uid {
		log.Printf("List with ID %v does not belongs to user with ID %v", list.ID, uid)
		c.Status(400)
		return
	}

	item := new(models.Item)
	item.Name = name
	item.Count = input.Count
	item.ListID = input.ListID

	if err := services.InsertItem(item); err != nil {
		log.Println("Failed to insert item", err)
		c.Status(500)
		return
	}

	c.JSON(item)
}

func UpdateItem(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID
	id := c.Params("ID")

	var input UpdateItemInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input", err)
		c.Status(400)
		return
	}

	name := strings.TrimSpace(input.Name)
	if input.Count <= 0 {
		log.Println("Receive count less than or equal to zero")
		c.Status(400)
		return
	}

	existingItem, err := services.FindItemByID(id)
	if err != nil {
		log.Println("Failed to find item by ID", err)
		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	list, err := services.GetListByID(existingItem.ListID)
	if err != nil {
		log.Println("Failed to find list by ID", err)
		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if list.UserID != uid {
		log.Printf("Item with ID %v does not belongs to user with ID %v", existingItem.ID, uid)
		c.Status(400)
		return
	}

	var item = models.Item{Name: name, Count: input.Count, Done: input.Done, Removed: input.Removed}
	item.ID = id

	updatedItem, err := services.UpdateItem(item)
	if err != nil {
		log.Println("Failed to update item", err)
		c.Status(500)
		return
	}

	c.JSON(updatedItem)
}
