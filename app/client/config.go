package main

import (
	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"golang.org/x/term"
)

type config struct {
	User      apicaller.User
	Term      *term.Terminal
	termState *term.State
	Window    int
	Client    apicaller.Client
}
