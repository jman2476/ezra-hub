package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jman2476/ezra-hub/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)
	}

	port := os.Getenv("SERVER_PORT")
	const filepathRoot = "."

	apiCfg := apiConfig{
		db: database.New(db),
	}
	log.Printf("apiCfg var: %v", apiCfg)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
