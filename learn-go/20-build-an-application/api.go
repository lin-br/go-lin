package api

import (
	"fmt"
	"net/http"
)

func PlayerServer(responseWriter http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, "20")
}
