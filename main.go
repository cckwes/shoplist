package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	auth "github.com/cckwes/shoplist/auth"
	database "github.com/cckwes/shoplist/db"
)

func main() {
	err := database.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}
	defer database.Close()

	app := fiber.New()

	app.Use(auth.JwtMiddleware)

	app.Get("/", func(context *fiber.Ctx) {
		userEmail := context.Locals("user").(map[string]string)["email"]

		context.JSON(&fiber.Map{
			"message": fmt.Sprintf("hello %v", userEmail),
		})
	})

	app.Listen(3000)
}
