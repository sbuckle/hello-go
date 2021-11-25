package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
