package main

import "fmt"

func commandHelp(cfg *config) error {
	for _, cmd := range getCommands() {
		fmt.Printf("\r%s:\n\r_____%s\n", cmd.name, cmd.description)
	}

	return nil
}
