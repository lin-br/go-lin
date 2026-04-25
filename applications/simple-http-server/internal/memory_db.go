package poker

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

func (i *InMemoryPlayerStore) GetLeagueTable() League {
	league := make([]Player, 0)
	for name, wins := range i.store {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
