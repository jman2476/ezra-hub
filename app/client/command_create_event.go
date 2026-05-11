package main

import (
	"fmt"
	"strconv"
	"time"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandCreateEvent(cfg *config) error {
	fmt.Println("Create a new event\n\r")
	var newEvent apicaller.NewEvent

	cfg.Term.SetPrompt("Title: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.Name = name

	cfg.Term.SetPrompt("Type: ")
	eventType, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.Type = eventType

	cfg.Term.SetPrompt("When is the event: YYYY-MM-DD hh:mm  ")
	date, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.OccursOn, err = time.Parse(time.DateTime, date+":00")
	if err != nil {
		return err
	}

	cfg.Term.SetPrompt("How many days is it valid for: ")
	duration, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	durNum, err := strconv.Atoi(duration)
	if err != nil {
		return err
	}
	durStr := strconv.Itoa(durNum*24) + "h"
	durParsed, err := time.ParseDuration(durStr)
	if err != nil {
		return nil
	}
	newEvent.ExpiresAt = newEvent.OccursOn.Add(durParsed)

	event, err := cfg.Client.NewEvent(newEvent)
	if err != nil {
		return err
	}
	printEvent(event)

	return nil
}

func printEvent(event apicaller.Event) {
	fmt.Printf("\rEvent name: %s\n", event.Name)
	fmt.Printf("\rCreated At: %v\n", event.CreatedAt)
	fmt.Printf("\rUpdated At: %v\n", event.UpdatedAt)
	fmt.Printf("\rType: %s\n", event.Type)
	fmt.Printf("\rOccurs On: %v\n", event.OccursOn)
	fmt.Printf("\rExpires At: %v\n", event.ExpiresAt)
}
