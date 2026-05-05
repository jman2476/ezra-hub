package main

import "golang.org/x/term"

type config struct {
	User      string
	Term      *term.Terminal
	termState *term.State
	Window    int
}
