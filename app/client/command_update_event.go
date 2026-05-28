package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"github.com/jman2476/ezra-hub/app/client/internal/menu"
)

func commandUpdateEvent(cfg *config) error {
	// first get user's created events
	events, err := cfg.Client.GetUserCreatedEvents(false)
	if err != nil {
		if err == fmt.Errorf("User owns no current events") {
			return fmt.Errorf("\rUser has no events to modify. Please create an event before you can modify it\r")
		}
		return fmt.Errorf("\rError getting events created by user: %w", err)
	}

	selected, err := menu.ItemMenu(events)
	if err != nil {
		return fmt.Errorf("Error selecting event from menu: %w", err)
	}

	fmt.Println("\rCurrent event attributes: ")
	printEvent(selected)

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
		eventUpdate.Name = selected.Name
	}

	cfg.Term.SetPrompt("Type: ")

	cfg.Term.SetPrompt("When is the event: YYYY-MM-DD hh:mm >")
	cfg.Term.SetPrompt("How many days is it valid for: ")
	cfg.Term.SetPrompt("Give a brief description of the event:\r\n")
	cfg.Term.SetPrompt("How many volunteers do you need? Type 0 if irrelavent. Invalid entries will be set to 0\n\rMin: ")
	cfg.Term.SetPrompt("Max: ")
	return nil
}
