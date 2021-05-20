package main

import (
	"log"
	"net/http"
)

const port = ":8090"

func main() {
	// Setup handlers...
	http.HandleFunc("/health-check", HealthCheckHandler)
	http.HandleFunc("/hello", HelloWorldHandler)

	// Listen server...
	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
