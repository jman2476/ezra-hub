package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func commandLogin(cfg *config) error {
	fmt.Println("Log in to your account\n\r")
	var loginData apicaller.UserLogin
	var user apicaller.User

	cfg.Term.SetPrompt("Name: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	loginData.Name = name

	cfg.Term.SetPrompt("Email: ")
	email, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	loginData.Email = email

	password, err := cfg.Term.ReadPassword("Password: ")
	if err != nil {
		return err
	}
	loginData.Password = password

	user, err = cfg.Client.LoginUser(loginData)
	if err != nil {
		return err
	}
	cfg.User = user
	printUser(user)

	err = msgbroker.SubscribeJSON(
		cfg.Connection,
		routing.ExchangeEzraDirect,
		routing.ActiveUserKey+"."+user.Name,
		routing.ActiveUserKey,
		msgbroker.SimpleQueueTransient,
		handlerActiveUsers(cfg),
	)
	// User subscriptions not properly getting sent from the server
	fmt.Printf("User subs: %v", user.Subscriptions)

	errSlice := resubQueues(cfg)
	for _, e := range errSlice {
		fmt.Println(
			fmt.Errorf("Err subscribing to %w", e),
		)
	}
	return nil
}

func resubQueues(cfg *config) (errSlice []error) {
	for _, cat := range cfg.User.Subscriptions {
		fmt.Printf("\rSubscribing to %s", cat)
		err := msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			cat+"."+cfg.User.Name,
			cat+".*",
			msgbroker.SimpleQueueTransient,
			handlerEvent(cfg),
		)
		errSlice = append(errSlice, fmt.Errorf("%s: %w", cat, err))
	}
	return
}
