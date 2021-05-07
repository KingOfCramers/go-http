package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
