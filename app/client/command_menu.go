package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/menu"
)

func commandMenu(cfg *config) error {
	command, _, err := menu.MenuRepl(listCommands(), 0)
	if err != nil {
		return fmt.Errorf("List command menu error: %w", err)
	}

	return getCommands()[command].callback(cfg)
}
