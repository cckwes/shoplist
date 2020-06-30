package server

import (
	"github.com/cckwes/shoplist/auth"
	"github.com/cckwes/shoplist/controllers"
	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/v1")

	api.Get("/lists", controllers.GetLists)
	api.Post("/lists", controllers.CreateList)
	api.Put("/lists/:ID", controllers.UpdateList)
	api.Get("/lists/:ID/items", controllers.GetItemsInList)

	api.Post("/items", controllers.CreateItem)
	api.Put("/items/:ID", controllers.UpdateItem)
}

func NewApp() *fiber.App {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}

	db.DB.LogMode(true)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.List{})
	db.DB.AutoMigrate(&models.Item{})

	app := fiber.New()

	app.Use(cors.New())
	app.Use(auth.JwtMiddleware)

	setupRoutes(app)

	return app
}
