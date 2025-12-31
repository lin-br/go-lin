package db

import (
	"encoding/json"
	"io"
	"os"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

type FileSystemPlayerStore struct {
	// update the database to *json.Encoder to don't
	// need to create a new encoder every time
	database *json.Encoder
	// Cache the league in memory so we don't need
	// to read from the file on every request.
	league model.League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	_, _ = file.Seek(0, io.SeekStart)
	league, _ := model.NewLeague(file)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

func (fs *FileSystemPlayerStore) GetLeagueTable() model.League {
	return fs.league
}

func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := fs.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (fs *FileSystemPlayerStore) RecordWin(name string) {
	player := fs.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		fs.league = append(fs.league, model.Player{Name: name, Wins: 1})
	}

	_ = fs.database.Encode(fs.league)
}
