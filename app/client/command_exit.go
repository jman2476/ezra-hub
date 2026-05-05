package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func commandExit(cfg *config) error {
	fmt.Println("\rClosing Ezra Hub client\n\r")
	term.Restore(cfg.Window, cfg.termState)
	os.Exit(0)
	return nil
}
