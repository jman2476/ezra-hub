package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	const filepathRoot = "."

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
