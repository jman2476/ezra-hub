package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
	"github.com/jman2476/ezra-hub/app/client/internal/msgbroker"
	"github.com/jman2476/ezra-hub/pkg/routing"
)

func commandLogin(cfg *config) error {
	fmt.Println("\rLog in to your account\n\r")
	var loginData apicaller.UserLogin
	var user apicaller.User

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
	fmt.Printf("\rUser subs: %v\n", user.Subscriptions)

	errSlice := resubQueues(cfg)
	for _, e := range errSlice {
		fmt.Println(
			fmt.Errorf("\rErr subscribing to %w\n", e),
		)
	}

	err = getSubbedEvents(cfg)
	if err != nil {
		fmt.Println("\rError getting subbed events: ", err)
	}

	return nil
}

func resubQueues(cfg *config) (errSlice []error) {
	for _, cat := range cfg.User.Subscriptions {
		fmt.Printf("\rSubscribing to new %s\n", cat)
		err := msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			cat+".new."+cfg.User.Name,
			cat+".new.*",
			msgbroker.SimpleQueueTransient,
			handlerEventNew(cfg),
		)
		if err != nil {
			errSlice = append(errSlice, fmt.Errorf("%s: %w", cat, err))
		}

		fmt.Printf("\rSubscribing to update %s\n", cat)
		err = msgbroker.SubscribeJSON(
			cfg.Connection,
			routing.ExchangeEzraTopic,
			cat+".update."+cfg.User.Name,
			cat+".update.*",
			msgbroker.SimpleQueueTransient,
			handlerEventUpdate(cfg),
		)
		if err != nil {
			errSlice = append(errSlice, fmt.Errorf("%s: %w", cat, err))
		}
	}

	return
}

func getSubbedEvents(cfg *config) error {
	events, err := cfg.Client.GetUserEvents(
		cfg.User.Subscriptions, false,
	)
	if err != nil {
		return err
	}

	for _, e := range events {
		cfg.Events[e.ID] = e
	}
	return nil
}
