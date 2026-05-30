package main

import (
	"database/sql"
	"fmt"
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
	args     []string
}

func main() {
	godotenv.Load("./app/server/.env")
	dbURL := os.Getenv("DB_URL")
	dbBackup := os.Getenv("DB_URL_BACKUP")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	rabbitConnString := os.Getenv("RABBIT_SERVER")
	rabbitBackupConn := os.Getenv("RABBIT_SERVER_DOCKER")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)

		db, err = sql.Open("postgres", dbBackup)
		if err != nil {
			log.Printf("Error opening database backup connection: %s", err)
		}
	}

	connection, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Printf("Error establishing connection to RabbitMQ server: %s", err)
		connection, err = amqp.Dial(rabbitBackupConn)
		if err != nil {
			log.Printf("Error using backup RabbitMQ connection: %s", err)
		}
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

	err = declareExchanges(channel)
	if err != nil {
		log.Printf("Error binding exchanges: %s", err)
	}

	port := os.Getenv("SERVER_PORT")
	const filepathRoot = "."

	apiCfg := apiConfig{
		db:       database.New(db),
		platform: platform,
		secret:   secret,
		channel:  channel,
	}

	if platform == "dev" {
		apiCfg.args = os.Args[1:]
	}

	mux := http.NewServeMux()

	// Development/Admin endpoints: Reset users
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetDB)
	mux.HandleFunc("GET /admin/users", apiCfg.handlerGetUsers)

	// User endpoints: Create, Login, Update
	mux.HandleFunc("POST /api/users", apiCfg.handlerNewUser)
	mux.HandleFunc("POST /api/login", apiCfg.handerLogIn)
	mux.HandleFunc("PATCH /api/users", apiCfg.authorize(apiCfg.handlerUpdateUser))
	mux.HandleFunc("PATCH /api/users/subs", apiCfg.authorize(apiCfg.handlerSubscribe))

	// Event endpoints: Create, Update
	mux.HandleFunc("POST /api/events", apiCfg.authorize(apiCfg.handlerNewEvent))
	mux.HandleFunc("PATCH /api/events/{id}", apiCfg.authorize(apiCfg.handlerRespondEvent))
	mux.HandleFunc("GET /api/events", apiCfg.authorize(apiCfg.handlerGetEventsByType))
	mux.HandleFunc("GET /api/events/users", apiCfg.authorize(apiCfg.handlerGetEventsbyUser))
	mux.HandleFunc("PUT /api/events/{id}", apiCfg.authorize(apiCfg.handlerUpdateEvent))

	// Token endpoints:
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("DEBUG: Addr string [%s] (Length: %d) (Bytes: %v)\n", server.Addr, len(server.Addr), []byte(server.Addr))

	log.Printf("Starting Ezra Hub server on port %s", port)
	log.Fatal(server.ListenAndServe())
}

func declareExchanges(ch *amqp.Channel) (err error) {
	err = ch.ExchangeDeclare(
		"ezra_direct", "direct",
		true, false, false, false,
		amqp.Table{},
	)
	if err != nil {
		return fmt.Errorf("Declare ezra_direct exchange error: %w", err)
	}

	err = ch.ExchangeDeclare(
		"ezra_topic", "topic",
		true, false, false, false,
		amqp.Table{},
	)
	if err != nil {
		return fmt.Errorf("Declare ezra_topic exchange error: %w", err)
	}

	return
}
