package simple_http_server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/lin-br/go-lin/applications/simple-http-server"
	"github.com/lin-br/go-lin/applications/simple-http-server/db"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := db.NewInMemoryPlayerStore()
	server := api.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
