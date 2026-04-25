package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lin-br/go-lin/applications/simple-http-server/internal"
)

const (
	DbFileName     = "game.db.json"
	FilePermission = 0666
)

func main() {
	// Open, or create, a file in the project root directory
	file, err := os.OpenFile(DbFileName, os.O_RDWR|os.O_CREATE, FilePermission)
	if err != nil {
		log.Fatalf("problem to open the file %s %v", DbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
