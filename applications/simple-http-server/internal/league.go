package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(reader io.Reader) (League, error) {
	league := make(League, 0)
	err := json.NewDecoder(reader).Decode(&league)

	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, nil
}

func (l League) Find(name string) *Player {
	for index, player := range l {
		if player.Name == name {
			return &l[index]
		}
	}
	return nil
}
