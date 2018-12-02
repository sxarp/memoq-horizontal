package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	fmt.Println("started server")

	server.ListenAndServe()
}
