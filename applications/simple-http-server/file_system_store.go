package simple_http_server

import (
	"encoding/json"
	"io"

	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

type FileSystemPlayerStore struct {
	// update the reader, and now we can read one more time
	// with the ReadWrite, now we can write too
	database io.ReadWriteSeeker
}

func (fs *FileSystemPlayerStore) GetLeague() model.League {
	_, _ = fs.database.Seek(0, io.SeekStart)
	league, _ := model.NewLeague(fs.database)
	return league
}

func (fs *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := fs.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (fs *FileSystemPlayerStore) RecordWin(name string) {
	league := fs.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, model.Player{Name: name, Wins: 1})
	}

	_, _ = fs.database.Seek(0, io.SeekStart)
	_ = json.NewEncoder(fs.database).Encode(league)
}
