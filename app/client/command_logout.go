package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandLogout(cfg *config) error {
	oldUser := cfg.User.Name
	cfg.User = apicaller.User{}
	cfg.Client.ClearTokens()

	fmt.Printf("User %s is signed out\r\n", oldUser)
	return nil
}
