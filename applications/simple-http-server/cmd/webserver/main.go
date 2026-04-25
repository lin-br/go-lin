package main

import (
	"log"
	"net/http"

	"github.com/lin-br/go-lin/applications/simple-http-server/internal"
)

const (
	DbFileName     = "game.db.json"
	FilePermission = 0666
)

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(DbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
