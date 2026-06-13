package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func resubQueues(cfg *config) (errSlice []error) {
	for _, cat := range cfg.User.Subscriptions {
		fmt.Printf("\rSubscribing to new %s\n", cat)
		err := msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			cat+".new."+cfg.User.Name,
			cat+".new.*",
			msgbroker.SimpleQueueTransient,
			handlerEventNew(cfg),
		)
		if err != nil {
			errSlice = append(errSlice, fmt.Errorf("%s: %w", cat, err))
		}

		fmt.Printf("\rSubscribing to update %s\n", cat)
		err = msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			cat+".update."+cfg.User.Name,
			cat+".update.*",
			msgbroker.SimpleQueueTransient,
			handlerEventUpdate(cfg),
		)
		if err != nil {
			errSlice = append(errSlice, fmt.Errorf("%s: %w", cat, err))
		}
	}

	return
}

func getSubbedEvents(cfg *config) error {
	events, err := cfg.Client.GetUserEvents(
		cfg.User.Subscriptions, false,
	)
	if err != nil {
		return err
	}

	for _, e := range events {
		cfg.Events[e.ID] = e
	}
	return nil
}
