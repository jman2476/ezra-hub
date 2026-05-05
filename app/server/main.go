package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db       *database.Queries
	platform string
}

func main() {
	godotenv.Load("./app/server/.env")
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)
	}

	port := os.Getenv("SERVER_PORT")
	const filepathRoot = "."

	apiCfg := apiConfig{
		db:       database.New(db),
		platform: platform,
	}

	mux := http.NewServeMux()

	// Development endpoints: Reset users
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetDB)

	// User endpoints: Create, Login, Update
	mux.HandleFunc("POST /api/users", apiCfg.handlerNewUser)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
