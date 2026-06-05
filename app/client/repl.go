package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jman2476/ezra-hub/app/client/internal/menu"
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

	// login/signup/exit
	cfg.loginOptions()

	// main loop
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
			cfg.handleScreen()
			continue
		}

		commandName := cleanedInput[0]
		command, ok := getCommands()[commandName]
		if ok {
			cfg.handleScreen()

			err := command.callback(cfg)

			if err != nil {
				fmt.Println("\r", err)
			}
			continue
		} else {
			cfg.handleScreen()

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

func (cfg *config) termNewLine() {
	prompt := "Ezra:" + cfg.User.Name + ">"
	cfg.Term.Write([]byte(prompt))
}

func (cfg *config) loginOptions() {
	var options = []string{"signup", "login", "exit"}
	commandName, _, err := menu.MenuRepl(options, 0)
	if err != nil {
		fmt.Println(
			fmt.Errorf("Menu error: %w", err),
		)
	}

	command, ok := getCommands()[commandName]
	if ok {
		err := command.callback(cfg)

		if err != nil {
			fmt.Println("\r", err)
		}

	} else {
		fmt.Println("\rUnknown command")

	}

}

func (cfg *config) handleScreen() error {
	x, y, err := cfg.getCursorPosition()
	if err != nil {
		return fmt.Errorf("Couldn't get cursor position: %w", err)
	}
	err = menu.ClearWindow()
	if err != nil {
		return fmt.Errorf("Error clearing window: %w", err)
	}
	cfg.drawAlertsBar()
	err = cfg.setAlert([]byte(""), x, y)
	if err != nil {
		return fmt.Errorf("Error drawing alerts bar: %w", err)
	}
	return nil
}

func (cfg *config) drawAlertsBar() {
	width, _, err := term.GetSize(int(cfg.Window))
	if err != nil {
		fmt.Println("Error getting terminal size")
		width = 10
	}
	var alertBar = make([]string, 3)
	alertBar[0] = "Alerts"
	for i := 0; i < width; i++ {
		if i > 5 {
			alertBar[0] += "*"
		}
		alertBar[2] += "*"
	}

	for i := range alertBar {
		fmt.Printf("\r%s\n\r", alertBar[i])
	}
}

func (cfg *config) getCursorPosition() (x, y int, err error) {
	_, err = os.Stdout.Write([]byte("\x1b[6n"))
	if err != nil {
		return -1, -1, err
	}

	var buf [32]byte
	n, err := os.Stdin.Read(buf[:])
	if err != nil {
		return
	}

	_, err = fmt.Sscanf(string(buf[:n]), "\x1b[%d;%dR", &x, &y)
	if err != nil {
		return -1, -1, err
	}

	return
}

func (cfg *config) setAlert(msg []byte, x, y int) error {
	w, _, err := term.GetSize(cfg.Window)
	var moveCursor strings.Builder
	var returnCursor strings.Builder
	var alert strings.Builder
	_, err = moveCursor.WriteString("\r")
	if err != nil {
		return fmt.Errorf("write moveCursor error: %w", err)
	}
	_, err = returnCursor.WriteString("\r")
	if err != nil {
		return fmt.Errorf("write returnCursor error: %w", err)
	}

	for range y + 1 {
		_, err = moveCursor.WriteString("\033[A")
		if err != nil {
			return fmt.Errorf("write moveCursor error: %w", err)
		}
		_, err = returnCursor.WriteString("\033[B")
		if err != nil {
			return fmt.Errorf("write returnCursor error: %w", err)
		}
	}
	for range x {
		_, err = returnCursor.WriteString("\033[C")
		if err != nil {
			return fmt.Errorf("write returnCursor error: %w", err)
		}
	}
	for i := range w {
		if len(msg) <= i {
			break
		}
		err = alert.WriteByte(msg[i])
		if err != nil {
			return fmt.Errorf("write alert error: %w", err)
		}
	}

	fmt.Print(moveCursor.String(), alert.String(), returnCursor.String())
	return nil
}
