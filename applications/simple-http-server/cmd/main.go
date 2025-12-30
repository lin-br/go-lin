package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/lin-br/go-lin/applications/simple-http-server"
)

const DbFileName = "game.db.json"

func main() {
	// Open, or create, a file in the project root directory
	file, err := os.OpenFile(DbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem to open the file %s %v", DbFileName, err)
	}

	store := api.NewFileSystemPlayerStore(file)
	server := api.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
