package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/menu"
)

func commandMenu(cfg *config) error {
	var cmdList []string
	for key, _ := range getCommands() {
		cmdList = append(cmdList, key)
	}
	command, _, err := menu.MenuRepl(cmdList, 0)
	if err != nil {
		return fmt.Errorf("List command menu error: %w", err)
	}

	return getCommands()[command].callback(cfg)
}
