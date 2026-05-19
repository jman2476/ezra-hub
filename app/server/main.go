package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type apiConfig struct {
	db       *database.Queries
	platform string
	secret   string
	channel  *amqp.Channel
}

func main() {
	godotenv.Load("./app/server/.env")
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	rabbitConnString := os.Getenv("RABBIT_SERVER")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)
	}

	connection, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Printf("Error establishing connection to RabbitMQ server: %s", err)
	} else {
		log.Printf("Connection to RabbitMQ server successful")
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Printf("Error creating channel: %s", err)
	} else {
		log.Printf("Successfully created channel %v", channel)
	}

	port := os.Getenv("SERVER_PORT")
	const filepathRoot = "."

	apiCfg := apiConfig{
		db:       database.New(db),
		platform: platform,
		secret:   secret,
		channel:  channel,
	}

	mux := http.NewServeMux()

	// Development/Admin endpoints: Reset users
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetDB)
	mux.HandleFunc("GET /admin/users", apiCfg.handlerGetUsers)

	// User endpoints: Create, Login, Update
	mux.HandleFunc("POST /api/users", apiCfg.handlerNewUser)
	mux.HandleFunc("POST /api/login", apiCfg.handerLogIn)
	mux.HandleFunc("PATCH /api/users", apiCfg.authorize(apiCfg.handlerSubscribe))

	// Event endpoints: Create, Update
	mux.HandleFunc("POST /api/events", apiCfg.authorize(apiCfg.handlerNewEvent))
	mux.HandleFunc("PATCH /api/events/{id}", apiCfg.authorize(apiCfg.handlerRespondEvent))

	// Token endpoints:
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting Ezra Hub server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
