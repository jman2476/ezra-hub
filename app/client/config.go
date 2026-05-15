package main

import (
	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/term"
)

type config struct {
	User       apicaller.User
	Term       *term.Terminal
	termState  *term.State
	Window     int
	Client     apicaller.Client
	Connection *amqp.Connection
	Events     []apicaller.Event
}
