package main

import "fmt"

func commandAlert(cfg *config) error {
	cfg.Term.SetPrompt("\rNew alert> ")
	input, err := cfg.Term.ReadLine()
	if err != nil {
		return fmt.Errorf("Readline error: %w", err)
	}
	x, y, err := cfg.getCursorPosition()
	if err != nil {
		return fmt.Errorf("Get cursor position error: %w", err)
	}
	err = cfg.setAlert([]byte(input), x, y)
	if err != nil {
		return fmt.Errorf("Set alert error: %w", err)
	}

	return nil
}
