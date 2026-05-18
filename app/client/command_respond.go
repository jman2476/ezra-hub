package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/menu"
)

func commandRespond(cfg *config) error {
	var listItems []string
	for _, e := range cfg.Events {
		listItems = append(listItems, e.Name)
	}
	listItems = append(listItems, "Return")

	_, index, err := menu.MenuRepl(listItems, 0)
	if err != nil {
		return fmt.Errorf("List event menu error: %w", err)
	}

	if index < len(cfg.Events) {
		printEvent(cfg.Events[index])
		fmt.Print("\n\rAre you available to go? ")
		going := menu.SelectYesNo()

		err := cfg.Client.RespondtoEvent(
			cfg.Events[index].ID,
			going,
		)
		if err != nil {
			return fmt.Errorf("Error responding to event: %w", err)
		}
	}

	return nil
}
