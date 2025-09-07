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

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []api.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
