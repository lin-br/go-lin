package db

import (
	"encoding/json"
	"io"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

type FileSystemPlayerStore struct {
	// update the reader, and now we can read one more time
	// with the ReadWrite, now we can write too
	database io.ReadWriteSeeker
	// Cache the league in memory so we don't need
	// to read from the file on every request.
	league model.League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	_, _ = database.Seek(0, io.SeekStart)
	league, _ := model.NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
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

	_, _ = fs.database.Seek(0, io.SeekStart)
	_ = json.NewEncoder(fs.database).Encode(fs.league)
}
