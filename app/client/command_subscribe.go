package main

import (
	"fmt"
	"strings"

	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func commandSubscribe(cfg *config) error {
	fmt.Println("\r\nWhat types of events do you want to get updates on?")
	fmt.Println("\rOptions: ride, shopping, check-in, meal, gathering, other")
	fmt.Println("\rType all your choices, seperated by commas\r")

	cfg.Term.SetPrompt("Event choices> ")
	input, err := cfg.Term.ReadLine()
	if err != nil {
		return fmt.Errorf("ReadLine error: %w", err)
	}

	choices := strings.Split(strings.ToLower(input), ",")
	for i, choice := range choices {
		choices[i] = strings.TrimSpace(choice)
	}

	var subMap = make(map[string]int)
	for _, choice := range choices {
		subMap[choice] = 1
	}
	fmt.Printf("\r\nSub Map: %v\r\n", subMap)

	newSubs, err := cfg.Client.SetSubscriptions(subMap, false)
	if err != nil {
		return fmt.Errorf("Set Subscriptions error: %w", err)
	}

	for _, sub := range newSubs {
		cfg.User.Subscriptions = append(cfg.User.Subscriptions, sub)
		err := msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			sub+"."+cfg.User.Name,
			sub+".*",
			msgbroker.SimpleQueueTransient,
			handlerEvent(cfg),
		)
		if err != nil {
			fmt.Println(fmt.Errorf("\r\nError subscribing to event type %s: %w", sub, err))
		} else {
			fmt.Printf("\rSuccessfully subscribed to the %s feed\n", sub)
		}
	}

	return nil
}
