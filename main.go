package main

import (
	"fmt"

	"github.com/gofiber/fiber"

	auth "github.com/cckwes/shoplist/auth"
)

func main() {
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
