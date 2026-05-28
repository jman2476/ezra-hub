package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	var eventUpdate apicaller.EventUpdate
	eventUpdate.OldType = selected.Category

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
	eventCategory, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if eventCategory != "" {
		eventUpdate.Category = strings.ToLower(eventCategory)
	} else {
		eventUpdate.Category = selected.Category
	}

	cfg.Term.SetPrompt("When is the event: YYYY-MM-DD hh:mm >")
	date, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if date != "" {
		eventUpdate.OccursOn, err = time.Parse(time.DateTime, date+":00")
		if err != nil {
			return err
		}
	} else {
		eventUpdate.OccursOn = selected.OccursOn
	}

	cfg.Term.SetPrompt("How many days is it valid for: ")
	duration, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if duration != "" {
		durNum, err := strconv.Atoi(duration)
		if err != nil {
			return err
		}
		durStr := strconv.Itoa(durNum*24) + "h"
		durParsed, err := time.ParseDuration(durStr)
		if err != nil {
			return err
		}
		eventUpdate.ExpiresAt = eventUpdate.OccursOn.Add(durParsed)
	} else {
		diff := selected.ExpiresAt.Sub(selected.OccursOn)
		eventUpdate.ExpiresAt = eventUpdate.OccursOn.Add(diff)
	}

	cfg.Term.SetPrompt("Give a brief description of the event:\r\n")
	description, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if description != "" {
		eventUpdate.Description = description
	} else {
		eventUpdate.Description = selected.Description
	}

	cfg.Term.SetPrompt("How many volunteers do you need? Type 0 if irrelavent. Invalid entries will be set to 0\n\rMin: ")
	min, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if min != "" {
		minInt, _ := strconv.Atoi(min) // ignore error, because error will set minInt to 0 anyway, which is what we want
		eventUpdate.MinVolunteers = int32(minInt)
	} else {
		eventUpdate.MinVolunteers = selected.MinVolunteers.Int32
	}

	cfg.Term.SetPrompt("Max: ")
	max, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if max != "" {
		maxInt, _ := strconv.Atoi(max) // ignore error, because error will set maxInt to 0 anyway, which is what we want
		eventUpdate.MaxVolunteers = int32(maxInt)
	} else {
		eventUpdate.MaxVolunteers = selected.MaxVolunteers.Int32
	}

	event, err := cfg.Client.UpdateEvent(eventUpdate, selected.ID)
	if err != nil {
		return err
	}
	printEvent(event)

	return nil
}
