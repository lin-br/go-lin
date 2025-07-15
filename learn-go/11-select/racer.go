package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
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
