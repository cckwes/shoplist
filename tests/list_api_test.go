package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/server"
	"github.com/cckwes/shoplist/services"
)

func TestGetList(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("List APIs", func() {
		g.Before(func() {
			SetupTests()
		})

		g.After(func() {
			TearDownTests()
		})

		g.BeforeEach(func() {
			ClearDB()
			CreateFixtures()
		})

		g.Describe("GET /v1/lists", func() {
			g.It("should return empty list", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get("/v1/lists").
					Header("Authorization", Bearer).
					Expect(t).
					Assert(jsonpath.Len(`$.lists`, 0)).
					Status(http.StatusOK).
					End()
			})

			g.It("should return non empty list", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get("/v1/lists").
					Header("Authorization", Bearer).
					Expect(t).
					Assert(jsonpath.Len(`$.lists`, 1)).
					Assert(jsonpath.Len(`$.lists[0].items`, 0)).
					Status(http.StatusOK).
					End()
			})

			g.It("should return non empty items in list", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)
				var item = models.Item{Name: "item1", Count: 1, ListID: list.ID}
				services.InsertItem(&item)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get("/v1/lists").
					Header("Authorization", Bearer).
					Expect(t).
					Assert(jsonpath.Len(`$.lists`, 1)).
					Assert(jsonpath.Len(`$.lists[0].items`, 1)).
					Status(http.StatusOK).
					End()
			})
		})

		g.Describe("GET /v1/lists/:ID", func() {
			g.It("should return 400 if list is not found", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get("/v1/lists/xxxx-yyyy").
					Header("Authorization", Bearer).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 403 if trying to access list of other user", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get(fmt.Sprintf("/v1/lists/%v", list.ID)).
					Header("Authorization", Bearer2).
					Expect(t).
					Status(http.StatusForbidden).
					End()
			})

			g.It("should be able to get list with valid ID", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Get(fmt.Sprintf("/v1/lists/%v", list.ID)).
					Header("Authorization", Bearer).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, list.Name)).
					Status(http.StatusOK).
					End()
			})
		})

		g.Describe("POST /v1/lists", func() {
			g.It("should return 400 for invalid input", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/lists").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"hello": "world"}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 400 for name that only has space", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/lists").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": " "}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should be able to create list", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/lists").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "party night"}`).
					Expect(t).
					Status(http.StatusOK).
					End()

				lists := services.FindListsByUserID(UserId)
				assert.Equal(t, "party night", lists[0].Name)
			})
		})

		g.Describe("PUT /v1/lists/:ID", func() {
			g.It("should return 400 if list is not found", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put("/v1/lists/xxxx-yyyy").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "groceries"}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 403 if trying to update list that's not belongs to the user", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/lists/%v", list.ID)).
					Header("Authorization", Bearer2).
					Header("Content-type", "application/json").
					Body(`{"name": "groceries"}`).
					Expect(t).
					Status(http.StatusForbidden).
					End()
			})

			g.It("should be able to update list", func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/lists/%v", list.ID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "groceries"}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "groceries")).
					Status(http.StatusOK).
					End()
			})
		})
	})
}
