package db

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, cleanDatabaseFunc := createTempFile(t, "12345")
	defer cleanDatabaseFunc()

	tape := &tape{file}
	_, _ = tape.Write([]byte("abc"))
	_, _ = file.Seek(0, io.SeekStart)

	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
