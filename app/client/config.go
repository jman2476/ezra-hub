package main

import (
	"github.com/google/uuid"
	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/term"
)

type config struct {
	User            apicaller.User
	Term            *term.Terminal
	termState       *term.State
	Window          int
	Client          apicaller.Client
	Connection      *amqp.Connection
	Events          map[uuid.UUID]apicaller.Event
	MostRecentAlert string
}
