package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Ezra Hub client")
	ezraClient := apicaller.NewClient(300 * time.Second)

	godotenv.Load("./app/client/.env")
	rabbitConnString := os.Getenv("RABBIT_SERVER")

	connection, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ server: %s", err)
	}
	defer connection.Close()

	cfg := &config{
		Window:     int(os.Stdin.Fd()), // sets reference for term window
		Client:     ezraClient,
		Connection: connection,
		Events:     make(map[uuid.UUID]apicaller.Event),
	}

	startRepl(cfg)
}
