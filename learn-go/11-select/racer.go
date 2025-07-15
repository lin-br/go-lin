package main

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	start1 := time.Now()
	http.Get(url1)
	duration1 := time.Since(start1)

	start2 := time.Now()
	http.Get(url2)
	duration2 := time.Since(start2)

	if duration1 < duration2 {
		return url1
	}
	return url2
}
