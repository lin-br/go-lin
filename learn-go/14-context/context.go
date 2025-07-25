package main

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()
		select {
		case fetched := <-data:
			_, _ = fmt.Fprint(w, fetched)
		case <-context.Done():
			store.Cancel()
		}
	}
}
