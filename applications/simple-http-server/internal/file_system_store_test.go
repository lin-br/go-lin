package poker

import (
	"os"
	"testing"
)

func assertScoreEquals(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// With the change to ReadWriteSeeker in FileSystemPlayerStore.database,
// the strings.Reader can't work, so we create a temporary file.
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, _ = tempFile.Write([]byte(initialData))

	removeFile := func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabaseFunc := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabaseFunc()

		store, err := NewFileSystemPlayerStore(database)

		got := store.GetLeagueTable()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)
		AssertNoError(t, err)

		// If we read again, the test will not pass.
		got = store.GetLeagueTable()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabaseFunc := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabaseFunc()

		store, err := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
		AssertNoError(t, err)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
		AssertNoError(t, err)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
		AssertNoError(t, err)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})

	//file_system_store_test.go
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeagueTable()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeagueTable()
		AssertLeague(t, got, want)
	})
}
