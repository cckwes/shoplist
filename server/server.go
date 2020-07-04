package server

import (
	"github.com/cckwes/shoplist/auth"
	"github.com/cckwes/shoplist/controllers"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/v1")

	api.Get("/lists", controllers.GetLists)
	api.Get("/lists/:ID", controllers.GetList)
	api.Post("/lists", controllers.CreateList)
	api.Put("/lists/:ID", controllers.UpdateList)

	api.Post("/items", controllers.CreateItem)
	api.Put("/items/:ID", controllers.UpdateItem)
}

func NewApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(auth.JwtMiddleware)

	setupRoutes(app)

	return app
}
