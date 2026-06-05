package main

import (
	"slices"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
	menuOrder   int
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "List commands for Ezra Hub client",
			callback:    commandHelp,
			menuOrder:   0,
		},
		"exit": {
			name:        "exit",
			description: "Exit Ezra Hub",
			callback:    commandExit,
			menuOrder:   20,
		},
		"signup": {
			name:        "signup",
			description: "Sign up a new user",
			callback:    commandSignUp,
			menuOrder:   2,
		},
		"login": {
			name:        "login",
			description: "Log into account",
			callback:    commandLogin,
			menuOrder:   1,
		},
		"create": {
			name:        "create event",
			description: "Create new event",
			callback:    commandCreateEvent,
			menuOrder:   5,
		},
		"logout": {
			name:        "logout",
			description: "Log out user",
			callback:    commandLogout,
			menuOrder:   4,
		},
		"subscribe": {
			name:        "subscribe",
			description: "Subscribe to feeds",
			callback:    commandSubscribe,
			menuOrder:   6,
		},
		"events": {
			name:        "events",
			description: "list events",
			callback:    commandListEvents,
			menuOrder:   7,
		},
		"respond": {
			name:        "respond",
			description: "Respond whether you can volunteer for an event",
			callback:    commandRespond,
			menuOrder:   8,
		},
		"menu": {
			name:        "menu",
			description: "show menu of all commands",
			callback:    commandMenu,
			menuOrder:   3,
		},
		"update-user": {
			name:        "update user",
			description: "update user information",
			callback:    commandUpdateUser,
			menuOrder:   9,
		},
		"update-event": {
			name:        "update event",
			description: "update event information",
			callback:    commandUpdateEvent,
			menuOrder:   10,
		},
		// "alert": {
		// 	name:        "alert",
		// 	description: "tmp/set alert bar",
		// 	callback:    commandAlert,
		// 	menuOrder:   11,
		// },
	}
}

func listCommands() (list []string) {
	cmds := getCommands()
	list = make([]string, 0, len(cmds))
	for key := range cmds {
		list = append(list, key)
	}

	slices.SortFunc(list, func(a, b string) int {
		return cmds[a].menuOrder - cmds[b].menuOrder
	})
	return
}
