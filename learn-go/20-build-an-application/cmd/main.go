package main

import (
	"log"
	"net/http"

	api "github.com/lin-br/go-lin/learn-go/20-build-an-application"
)

func main() {
	handler := http.HandlerFunc(api.PlayerServer)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
