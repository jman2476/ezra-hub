package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

func startRepl(cfg *config) {
	oldSate, err := term.MakeRaw(cfg.Window)
	if err != nil {
		fmt.Println(
			fmt.Errorf("Initialization error: cannot get raw terminal input: %w", err),
		)
	}
	defer term.Restore(cfg.Window, oldSate)

	rw := setReaderWriter(os.Stdin, os.Stdout)
	cfg.Term = term.NewTerminal(rw, "")
	cfg.termState = oldSate

	for {
		prompt := "Ezra"
		if cfg.User.Name != "" {
			prompt = "Ezra:" + cfg.User.Name
		}
		cfg.Term.SetPrompt(fmt.Sprintf("\r%s > ", prompt))

		buffer, err := cfg.Term.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("\rEnd of file")
				commandExit(cfg)
			}
			fmt.Println("\r", err)
		}
		cleanedInput := cleanInput(buffer)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg)

			if err != nil {
				fmt.Println("\r", err)
			}
			continue
		} else {
			fmt.Println("\rUnknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}

func setReaderWriter(in, out *os.File) io.ReadWriter {
	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: in,
		Writer: out,
	}

	return rw
}
