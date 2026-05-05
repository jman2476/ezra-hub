package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit Ezra Hub",
			callback:    commandExit,
		},
		"signup": {
			name:        "signup",
			description: "Sign up a new user",
			callback:    commandSignUp,
		},
	}
}
