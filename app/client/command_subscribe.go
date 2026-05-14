package main

import (
	"fmt"
	"strings"
)

func commandSubscribe(cfg *config) error {
	fmt.Println("\r\nWhat types of events do you want to get updates on?")
	fmt.Println("\rOptions: ride, shopping, check-in, meal, gathering, other")
	fmt.Println("\rType all your choices, seperated by commas\r")

	cfg.Term.SetPrompt("Event choices> ")
	input, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}

	choices := strings.Split(strings.ToLower(input), ",")
	for i, choice := range choices {
		choices[i] = strings.TrimSpace(choice)
	}

	var subMap = make(map[string]int)
	for _, choice := range choices {
		subMap[choice] = 1
	}

	if cfg.User.Subscriptions != nil {
		for _, sub := range cfg.User.Subscriptions {
			subMap[sub] = 1
		}
	}

	return nil
}
