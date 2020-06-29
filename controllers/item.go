package controllers

import (
	"log"

	"github.com/gofiber/fiber"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/services"
)

type CreateItemInput struct {
	Name   string `json:"name"`
	Count  uint   `json:"count"`
	ListID string `json:"list_id"`
}

func CreateItem(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID

	var input CreateItemInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input ", err)
		c.Status(400)
		return
	}

	list, err := services.GetListByID(input.ListID)
	if err != nil {
		log.Println("Failed to find list ", err)
		c.Status(500)
		return
	}
	if list.UserID != uid {
		log.Printf("List with ID %v does not belongs to user with ID %v", list.ID, uid)
		c.Status(403)
		return
	}

	item := new(models.Item)
	item.Name = input.Name
	item.Count = input.Count
	item.ListID = input.ListID

	if err := services.InsertItem(item); err != nil {
		log.Println("Failed to insert item ", err)
		c.Status(500)
		return
	}

	c.JSON(item)
}

func UpdateItem(c *fiber.Ctx) {
}
