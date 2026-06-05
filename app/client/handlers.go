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

		alert := fmt.Sprintf("%s logged in", user.Name)
		cfg.MostRecentAlert = alert

		return msgbroker.Ack
	}
}

func handlerEventNew(cfg *config) func(apicaller.EventwName) msgbroker.AckType {
	return func(e apicaller.EventwName) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))
		defer cfg.termNewLine()

		alert := fmt.Sprintf("New %s Event Incoming from %s!", strings.ToUpper(e.Category), e.Creator)
		cfg.MostRecentAlert = alert

		cfg.Events[e.ID] = e.Event

		return msgbroker.Ack
	}
}

func handlerEventUpdate(cfg *config) func(apicaller.Event) msgbroker.AckType {
	return func(e apicaller.Event) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))
		defer cfg.termNewLine()

		alert := fmt.Sprintf("Updated %s Event Incoming!", strings.ToUpper(e.Category))
		cfg.MostRecentAlert = alert
		cfg.Events[e.ID] = e
		return msgbroker.Ack
	}
}
