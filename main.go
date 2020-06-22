package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	auth "github.com/cckwes/shoplist/auth"
)

func main() {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("Failed to connect to databasae")
	}
	defer db.Close()

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
