package main

import (
	"fmt"
	"strconv"
	"strings"
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
	eventCategory, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.Category = strings.ToLower(eventCategory)

	cfg.Term.SetPrompt("When is the event: YYYY-MM-DD hh:mm >")
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
		return err
	}
	newEvent.ExpiresAt = newEvent.OccursOn.Add(durParsed)

	cfg.Term.SetPrompt("Where is the event happening:\r\n")
	location, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.Location = location

	cfg.Term.SetPrompt("Give a brief description of the event:\r\n")
	description, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newEvent.Description = description

	cfg.Term.SetPrompt("How many volunteers do you need? Type 0 if irrelavent. \rInvalid entries will be set to 0\n\rMin: ")
	min, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	minInt, _ := strconv.Atoi(min) // ignore error, because error will set minInt to 0 anyway, which is what we want
	newEvent.MinVolunteers = int32(minInt)

	cfg.Term.SetPrompt("Max: ")
	max, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	maxInt, _ := strconv.Atoi(max) // ignore error, because error will set maxInt to 0 anyway, which is what we want
	newEvent.MaxVolunteers = int32(maxInt)

	event, err := cfg.Client.NewEvent(newEvent)
	if err != nil {
		return err
	}
	printEventwName(event)

	return nil
}

func printEvent(event apicaller.Event) {
	fmt.Printf("\rEvent name: %s\n", event.Name)
	fmt.Printf("\rCreated At: %v\n", event.CreatedAt)
	fmt.Printf("\rUpdated At: %v\n", event.UpdatedAt)
	fmt.Printf("\rCategory: %s\n", event.Category)
	fmt.Printf("\rOccurs On: %v\n", event.OccursOn)
	fmt.Printf("\rExpires At: %v\n", event.ExpiresAt)
	fmt.Printf("\rLocation: %s\n", event.Location)
	printVolunteersNeeded(event)
	printNumGoing(event)
	fmt.Printf("\rDescription: %v\n", event.Description)
}

func printEventwName(event apicaller.EventwName) {
	fmt.Printf("\rNew event from %s\n", event.Creator)
	printEvent(event.Event)
}

func printVolunteersNeeded(event apicaller.Event) {
	volunteer := "Volunteers: "
	if event.MinVolunteers.Valid {
		volunteer += fmt.Sprintf("Min: %d ", event.MinVolunteers.Int32)
	}
	if event.MaxVolunteers.Valid {
		volunteer += fmt.Sprintf("Max: %d", event.MaxVolunteers.Int32)
	}

	if !event.MinVolunteers.Valid && !event.MaxVolunteers.Valid {
		volunteer += "Any"
	}

	fmt.Printf("\r%s\n", volunteer)
}

func printNumGoing(event apicaller.Event) {
	count := len(event.Respondants)
	fmt.Printf("\rPeople going: %d\n", count)
}
