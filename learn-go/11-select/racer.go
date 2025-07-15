package main

import (
	"net/http"
)

func Racer(url1, url2 string) (winner string) {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan struct{} {
	channels := make(chan struct{})
	go func() {
		http.Get(url)
		close(channels)
	}()
	return channels
}
