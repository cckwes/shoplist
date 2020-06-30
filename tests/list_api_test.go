package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/steinfletcher/apitest"

	"github.com/cckwes/shoplist/server"
)

func TestMain(m *testing.M) {
	SetupTests()
	code := m.Run()
	TearDownTests()
	os.Exit(code)
}

func TestGetList_EmptyList(t *testing.T) {
	ClearDB()

	apitest.New().
		HandlerFunc(FiberToHandler(server.NewApp())).
		Get("/v1/lists").
		Header("Authorization", Bearer).
		Expect(t).
		Body("[]").
		Status(http.StatusOK).
		End()
}
