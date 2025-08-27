package main

import (
	"log"
	"net/http"

	api "github.com/lin-br/go-lin/learn-go/20-build-an-application"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &api.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}
