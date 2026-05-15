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
		"login": {
			name:        "login",
			description: "Log into account",
			callback:    commandLogin,
		},
		"create": {
			name:        "create event",
			description: "Create new event",
			callback:    commandCreateEvent,
		},
		"logout": {
			name:        "logout",
			description: "Log out user",
			callback:    commandLogout,
		},
		"subscribe": {
			name:        "subscribe",
			description: "Subscribe to feeds",
			callback:    commandSubscribe,
		},
		"events": {
			name:        "events",
			description: "list events",
			callback:    commandListEvents,
		},
	}
}
