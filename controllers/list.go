package controllers

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/services"
)

type CreateListInput struct {
	Name string `json:"name"`
}

func GetLists(c *fiber.Ctx) {
	id := c.Locals("user").(models.User).ID

	items := services.FindListsByUserID(id)
	c.JSON(items)
}

func CreateList(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID

	var input CreateListInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input ", err)
		c.Status(400)
		return
	}

	var list models.List
	list.Name = input.Name
	list.UserID = uid
	err := services.InsertList(&list)

	if err != nil {
		log.Println("Failed to create list ", err)
		c.Status(500)
		return
	}

	c.JSON(list)
}

func UpdateList(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID
	id := c.Params("ID")

	l, err := services.GetListByID(id)
	if err != nil {
		log.Println("Error when getting list ", err)

		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if l.UserID != uid {
		log.Panicln("Trying to update a list that's not belongs to the user")
		c.Status(403)
		return
	}

	var input CreateListInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input ", err)
		c.Status(400)
		return
	}

	var list models.List
	list.Name = input.Name
	list.ID = id
	updatedList, err := services.UpdateList(list)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(404)
			return
		}

		log.Println("Failed to update list ", err)
		c.Status(500)
		return
	}

	c.JSON(updatedList)
}

func GetItemsInList(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID
	id := c.Params("ID")

	l, err := services.GetListByID(id)
	if err != nil {
		log.Println("Error when getting list ", err)

		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if l.UserID != uid {
		log.Panicln("Trying to get items in a list that's not belongs to the user")
		c.Status(403)
		return
	}

	items, err := services.GetItemsByListID(id)
	if err != nil {
		log.Println("Error when getting items in list ", err)
		c.Status(500)
		return
	}

	c.JSON(items)
}
