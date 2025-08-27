package main

import (
	"log"
	"net/http"

	api "github.com/lin-br/go-lin/learn-go/20-build-an-application"
	"github.com/lin-br/go-lin/learn-go/20-build-an-application/db"
)

func main() {
	server := &api.PlayerServer{Store: db.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", server))
}
