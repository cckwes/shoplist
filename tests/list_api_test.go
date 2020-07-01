package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"github.com/cckwes/shoplist/models"
	"github.com/cckwes/shoplist/server"
	"github.com/cckwes/shoplist/services"
)

func TestMain(m *testing.M) {
	SetupTests()
	code := m.Run()
	TearDownTests()
	os.Exit(code)
}

func TestGetList_EmptyList(t *testing.T) {
	ClearDB()
	CreateFixtures()

	apitest.New().
		HandlerFunc(FiberToHandler(server.NewApp())).
		Get("/v1/lists").
		Header("Authorization", Bearer).
		Expect(t).
		Assert(jsonpath.Len(`$.lists`, 0)).
		Status(http.StatusOK).
		End()
}

func TestGetList_NonEmptyList(t *testing.T) {
	ClearDB()
	CreateFixtures()
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
}
