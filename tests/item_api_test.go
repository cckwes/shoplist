package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/server"
	"github.com/cckwes/shoplist/services"
	"github.com/franela/goblin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestItemAPI(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Item APIs", func() {
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

		g.Describe("POST /v1/items", func() {
			var listID string

			g.BeforeEach(func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)
				listID = list.ID
			})

			g.It("should return 400 for empty name", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/items").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(fmt.Sprintf(`{"name": "", "count": 1, "list_id": "%v"}`, listID)).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 400 for 0 count", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/items").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(fmt.Sprintf(`{"name": "Egg", "count": 0, "list_id": "%v"}`, listID)).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 400 if list id is invalid", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/items").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 1, "list_id": "invalid-list-id"}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 400 if list doesn't belongs to the user", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/items").
					Header("Authorization", Bearer2).
					Header("Content-type", "application/json").
					Body(fmt.Sprintf(`{"name": "Egg", "count": 1, "list_id": "%v"}`, listID)).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should be able to create item", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Post("/v1/items").
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(fmt.Sprintf(`{"name": "Egg", "count": 1, "list_id": "%v"}`, listID)).
					Expect(t).
					Status(http.StatusOK).
					Assert(jsonpath.Equal(`$.name`, "Egg")).
					End()
			})
		})

		g.Describe("PUT /v1/items/:ID", func() {
			var itemID string

			g.BeforeEach(func() {
				var list = models.List{Name: "default", UserID: UserId}
				services.InsertList(&list)

				var item = models.Item{Name: "Egg", Count: 2, ListID: list.ID, Done: true, Removed: false}
				services.InsertItem(&item)
				itemID = item.ID
			})

			g.It("should not update count if it's 0", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 0, "done": false, "removed": false}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should return 400 if item doesn't belongs to the user", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer2).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 1, "done": false, "removed": false}`).
					Expect(t).
					Status(http.StatusBadRequest).
					End()
			})

			g.It("should be able to update item name only", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Bread", "count": 2, "done": false, "removed": false}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "Bread")).
					Assert(jsonpath.Equal(`$.count`, float64(2))).
					Status(http.StatusOK).
					End()
			})

			g.It("should be able to update item count only", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 1, "done": false, "removed": false}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "Egg")).
					Assert(jsonpath.Equal(`$.count`, float64(1))).
					Status(http.StatusOK).
					End()
			})

			g.It("should be able to mark item as undone", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 2, "done": false, "removed": false}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "Egg")).
					Assert(jsonpath.Equal(`$.count`, float64(2))).
					Assert(jsonpath.Equal(`$.done`, false)).
					Assert(jsonpath.Equal(`$.removed`, false)).
					Status(http.StatusOK).
					End()
			})

			g.It("should be able to mark item as removed", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Egg", "count": 2, "done": false, "removed": true}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "Egg")).
					Assert(jsonpath.Equal(`$.count`, float64(2))).
					Assert(jsonpath.Equal(`$.done`, false)).
					Assert(jsonpath.Equal(`$.removed`, true)).
					Status(http.StatusOK).
					End()
			})

			g.It("should be able to update item name and count", func() {
				apitest.New().
					HandlerFunc(FiberToHandler(server.NewApp())).
					Put(fmt.Sprintf("/v1/items/%v", itemID)).
					Header("Authorization", Bearer).
					Header("Content-type", "application/json").
					Body(`{"name": "Bread", "count": 1, "done": false, "removed": true}`).
					Expect(t).
					Assert(jsonpath.Equal(`$.name`, "Bread")).
					Assert(jsonpath.Equal(`$.count`, float64(1))).
					Status(http.StatusOK).
					End()
			})
		})
	})
}
