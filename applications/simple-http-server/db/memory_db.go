package db

import api "github.com/lin-br/go-lin/applications/simple-http-server"

// NewInMemoryPlayerStore a constructor function
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

func (i *InMemoryPlayerStore) GetLeagueTable() []api.Player {
	var league []api.Player
	for name, wins := range i.store {
		league = append(league, api.Player{Name: name, Wins: wins})
	}
	return league
}
