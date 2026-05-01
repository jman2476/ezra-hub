package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "3294"
	const filepathRoot = "."

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port 3294")
	log.Fatal(server.ListenAndServe())
}
