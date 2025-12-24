package simple_http_server

import (
	"strings"
	"testing"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []model.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// If we read again, the test will not pass.
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
