package main

import (
	"fmt"
	"io"
	"os"

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
		if cfg.User != "" {
			prompt = "Ezra:" + cfg.User
		}
		cfg.Term.SetPrompt(fmt.Sprintf("\r%s > ", prompt))

		buffer, err := cfg.Term.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("\rEnd of file")
				commandExit(cfg)
			}
			fmt.Println(
				fmt.Errorf("\r, err"),
			)
		}

		fmt.Printf("\rLine read: %s\n", buffer)
	}
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
