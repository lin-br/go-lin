package db

import (
	"github.com/lin-br/go-lin/applications/simple-http-server/model"
)

// NewInMemoryPlayerStore used for testing purposes
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeagueTable() model.League {
	league := make([]model.Player, 0)
	for name, wins := range i.store {
		league = append(league, model.Player{Name: name, Wins: wins})
	}
	return league
}
