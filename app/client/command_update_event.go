package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandUpdateEvent(cfg *config) error {
	// first get choose an event from the the user's created events

	// set that equal to var currentEvent apicaller.Event
	var eventUpdate apicaller.NewEvent

	fmt.Println("\rType in the fields you want to update, leave the rest blank\r")

	cfg.Term.SetPrompt("Name: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if name != "" {
		eventUpdate.Name = name
	} else {
	}

	cfg.Term.SetPrompt("Name: ")

	cfg.Term.SetPrompt("Name: ")
	cfg.Term.SetPrompt("Name: ")
	cfg.Term.SetPrompt("Name: ")
	cfg.Term.SetPrompt("Name: ")
	return nil
}
