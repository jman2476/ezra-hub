package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/menu"
)

func commandListEvents(cfg *config) error {
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
	}

	return nil
}
