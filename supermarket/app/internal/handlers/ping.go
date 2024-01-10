// This script contains the handler for the second task in the API class in Go

package internal

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}
