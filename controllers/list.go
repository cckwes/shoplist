package controllers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/services"
)

type CreateListInput struct {
	Name string `json:"name"`
}

type Lists struct {
	Lists []models.List `json:"lists"`
}

func GetLists(c *fiber.Ctx) {
	id := c.Locals("user").(models.User).ID

	lists := services.FindListsByUserID(id)
	c.JSON(Lists{Lists: lists})
}

func CreateList(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID

	var input CreateListInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input", err)
		c.Status(400)
		return
	}

	name := strings.TrimSpace(input.Name)
	if len(name) == 0 {
		log.Println("Received invalid name for create list")
		c.Status(400)
		return
	}

	var list models.List
	list.Name = name
	list.UserID = uid
	err := services.InsertList(&list)

	if err != nil {
		log.Println("Failed to create list", err)
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
		log.Println("Error when getting list", err)

		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if l.UserID != uid {
		log.Println("Trying to update a list that's not belongs to the user")
		c.Status(400)
		return
	}

	var input CreateListInput
	if err := c.BodyParser(&input); err != nil {
		log.Println("Unable to parse input", err)
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

		log.Println("Failed to update list", err)
		c.Status(500)
		return
	}

	c.JSON(updatedList)
}

func GetList(c *fiber.Ctx) {
	uid := c.Locals("user").(models.User).ID
	id := c.Params("ID")

	list, err := services.GetListByID(id)
	if err != nil {
		log.Println("Error when getting list", err)

		if err == gorm.ErrRecordNotFound {
			c.Status(400)
		} else {
			c.Status(500)
		}
		return
	}
	if list.UserID != uid {
		log.Println("Trying to get items in a list that's not belongs to the user")
		c.Status(400)
		return
	}

	c.JSON(list)
}
