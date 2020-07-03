package tests

import (
	"io"
	"net/http"
	"os"

	"github.com/cckwes/shoplist/db"
	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/services"
	"github.com/gofiber/fiber"
)

const Bearer string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ.o88lC_FoBgf5Ke5IgezBPPHvBQ0n5jAsnm932jxPSMI"
const Bearer2 string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIn0.ZPlIU_AuF81AgchYWxoj9s3r2OXi8e_i8-wrp1JVLkc"

const email string = "test@example.com"
const email2 string = "test2@example.com"

var UserId string = ""
var UserId2 string = ""

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

func CreateFixtures() {
	user, err := services.GetOrCreateUser(email)
	if err != nil {
		panic("Failed to setup fixture")
	}
	UserId = user.ID

	user2, err := services.GetOrCreateUser(email2)
	if err != nil {
		panic("Failed to setup fixture")
	}
	UserId2 = user2.ID
}

func ClearDB() {
	db.DB.Exec("DELETE FROM items;")
	db.DB.Exec("DELETE FROM lists;")
	db.DB.Exec("DELETE FROM users;")
}

func SetupTests() {
	err := db.Open()
	if err != nil {
		panic("Failed to connect to databasae")
	}

	// db.DB.LogMode(false)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.List{})
	db.DB.AutoMigrate(&models.Item{})
	os.Setenv("JWT_DEV_SECRET", "ABC")
}

func TearDownTests() {
	db.Close()
}
