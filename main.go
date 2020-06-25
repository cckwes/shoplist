package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	auth "github.com/cckwes/shoplist/auth"
	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
)

func main() {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}
	defer db.Close()

	// db.DB.LogMode(true)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Item{})

	app := fiber.New()

	app.Use(auth.JwtMiddleware)

	app.Get("/", func(context *fiber.Ctx) {
		userEmail := context.Locals("user").(models.User).Email

		context.JSON(&fiber.Map{
			"message": fmt.Sprintf("hello %v", userEmail),
		})
	})

	app.Listen(3000)
}
