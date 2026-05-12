package main

import (
	"fmt"

	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func handlerActiveUsers(cfg *config) func(routing.ActiveUser) msgbroker.AckType {
	return func(user routing.ActiveUser) msgbroker.AckType {
		cfg.Term.Write([]byte("\r\n"))

		fmt.Println("\r***************************\r")
		fmt.Printf("----------%s logged in----------\r\n", user.Name)
		fmt.Println("***************************\r")

		cfg.termNewLine()
		return msgbroker.Ack
	}
}
