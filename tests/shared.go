package tests

import (
	"io"
	"net/http"
	"os"

	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
	"github.com/gofiber/fiber"
)

func FiberToHandler(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

func ClearDB() {
	db.DB.Exec("TRUNCATE TABLE items;")
	db.DB.Exec("TRUNCATE TABLE lists;")
	db.DB.Exec("TRUNCATE TABLE users;")
}

func SetupTests() {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}

	db.DB.LogMode(false)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.List{})
	db.DB.AutoMigrate(&models.Item{})
	os.Setenv("JWT_DEV_SECRET", "ABC")
}

func TearDownTests() {
	db.Close()
}

const Bearer string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ.o88lC_FoBgf5Ke5IgezBPPHvBQ0n5jAsnm932jxPSMI"
