package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Returns the string "Hello!"
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
