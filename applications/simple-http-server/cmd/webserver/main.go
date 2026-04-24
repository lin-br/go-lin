package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/lin-br/go-lin/applications/simple-http-server"
	"github.com/lin-br/go-lin/applications/simple-http-server/db"
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

	store, err := db.NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := api.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
