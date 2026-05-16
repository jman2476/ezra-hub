package msgbroker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jman2476/ezra-hub/app/server/internal/database"
	"github.com/jman2476/ezra-hub/app/server/internal/outgoing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T, db *database.Queries) error {
	// There's an error here, and I need to find it
	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("Marshaling error: %w", err)
	}

	pub := amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	}
	outgoingMessage := outgoing.Message{
		Exchange:   exchange,
		Key:        key,
		Publishing: pub,
	}

	id, err := outgoing.LogOutgoingMessage(db, outgoingMessage)
	if err != nil {
		return fmt.Errorf("Outbox log error: %w", err)
	}
	defer outgoing.LogMessageSent(db, id)

	return ch.PublishWithContext(
		context.Background(),
		exchange, key,
		false, false,
		pub,
	)
}
