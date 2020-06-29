package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/cckwes/shoplist/auth"
	"github.com/cckwes/shoplist/controllers"
	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
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

func newApp() *fiber.App {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}
	defer db.Close()

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

func main() {
	app := newApp()

	app.Listen(3000)
}
