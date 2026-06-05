package main

import (
	"fmt"
	"strings"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func handlerActiveUsers(cfg *config) func(routing.ActiveUser) msgbroker.AckType {
	return func(user routing.ActiveUser) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))
		defer cfg.termNewLine()

		fmt.Println("\r***************************\r")
		fmt.Printf("----------%s logged in----------\r\n", user.Name)
		fmt.Println("***************************\r")
		alert := fmt.Sprintf("%s logged in", user.Name)
		x, y, err := cfg.getCursorPosition()
		if err == nil {
			cfg.setAlert([]byte(alert), x, y)
		} else {
			cfg.setAlert([]byte(err.Error()), x, y)
		}

		return msgbroker.Ack
	}
}

func handlerEventNew(cfg *config) func(apicaller.EventwName) msgbroker.AckType {
	return func(e apicaller.EventwName) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))
		defer cfg.termNewLine()

		// fmt.Println("\r************************")
		// fmt.Printf("\rNew %s Event Incoming from %s!\n", strings.ToUpper(e.Category), e.Creator)
		// printEvent(e.Event)
		// fmt.Println("\r************************\r")
		alert := fmt.Sprintf("New %s Event Incoming from %s!", strings.ToUpper(e.Category), e.Creator)
		x, y, err := cfg.getCursorPosition()
		if err == nil {
			cfg.setAlert([]byte(alert), x, y)
		} else {
			cfg.setAlert([]byte(err.Error()), x, y)
		}

		cfg.Events[e.ID] = e.Event

		return msgbroker.Ack
	}
}

func handlerEventUpdate(cfg *config) func(apicaller.Event) msgbroker.AckType {
	return func(e apicaller.Event) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))
		defer cfg.termNewLine()

		// fmt.Println("\r************************")
		// fmt.Printf("\rUpdated %s Event Incoming!\n", strings.ToUpper(e.Category))
		// printEvent(e)
		// fmt.Println("\r************************\r")
		alert := fmt.Sprintf("Updated %s Event Incoming!", strings.ToUpper(e.Category))
		x, y, err := cfg.getCursorPosition()
		if err == nil {
			cfg.setAlert([]byte(alert), x, y)
		} else {
			cfg.setAlert([]byte(err.Error()), x, y)
		}
		cfg.Events[e.ID] = e

		return msgbroker.Ack
	}
}
