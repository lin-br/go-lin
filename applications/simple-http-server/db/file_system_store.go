package db

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

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

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := model.NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func (fs *FileSystemPlayerStore) GetLeagueTable() model.League {
	sort.Slice(fs.league, func(i, j int) bool {
		return fs.league[i].Wins > fs.league[j].Wins
	})
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

func initialisePlayerDBFile(file *os.File) error {
	_, _ = file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, _ = file.Write([]byte("[]"))
		_, _ = file.Seek(0, io.SeekStart)
	}

	return nil
}
