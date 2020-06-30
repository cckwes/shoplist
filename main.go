package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/server"
)

func main() {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}
	defer db.Close()

	db.DB.LogMode(true)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.List{})
	db.DB.AutoMigrate(&models.Item{})

	app := server.NewApp()

	app.Listen(3000)
}
