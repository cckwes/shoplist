package tests

import (
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/server"
	"github.com/cckwes/shoplist/services"
)

func TestGetList(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("GET /v1/lists", func() {
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
				Status(http.StatusOK).
				End()
		})
	})
}
