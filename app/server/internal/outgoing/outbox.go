package outgoing

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
	amqp "github.com/rabbitmq/amqp091-go"
)

func LogOutgoingMessage(db *database.Queries, msg Outgoing) (uuid.UUID, error) {
	data, err := EncodeGob(msg)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Log Outgoing Message error: %w", err)
	}

	out, err := db.CreateOutgoing(context.Background(), data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Create outgoing database entry error: %w", err)
	}

	return out.ID, nil
}

type Outgoing struct {
	Exchange   string
	Key        string
	Publishing amqp.Publishing
}

func LogMessageSent(db *database.Queries, id uuid.UUID) error {
	return db.UpdateByIDtoSENT(context.Background(), id)
}
