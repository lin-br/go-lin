package main

import (
	"log"
	"net/http"

	api "github.com/lin-br/go-lin/applications/simple-http-server"
	"github.com/lin-br/go-lin/applications/simple-http-server/db"
)

func main() {
	server := &api.PlayerServer{Store: db.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", server))
}
