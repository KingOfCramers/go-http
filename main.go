package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

const port = ":8090"

func main() {
	http.HandleFunc("/hello", hello)
	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
