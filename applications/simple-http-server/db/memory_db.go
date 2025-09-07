package db

import simple_http_server "github.com/lin-br/go-lin/applications/simple-http-server"

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

func (i *InMemoryPlayerStore) GetLeagueTable() []simple_http_server.Player {
	return nil
}
