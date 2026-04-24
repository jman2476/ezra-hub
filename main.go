package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":3294",
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port 3294")
	log.Fatal(server.ListenAndServe())
}
