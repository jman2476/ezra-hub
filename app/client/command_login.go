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

	user, err := cfg.Client.LoginUser(loginData)
	if err != nil {
		return err
	}
	cfg.User = user
	printUser(user)

	_, _, err = msgbroker.DeclareAndBind(
		cfg.Connection,
		routing.ExchangeEzraDirect,
		routing.ActiveUserKey+"."+user.Name,
		routing.ActiveUserKey,
		msgbroker.SimpleQueueTransient,
	)
	if err != nil {
		fmt.Printf("Errors encountered in declare and bind: %s\r\n", err)
	}

	err = msgbroker.SubscribeJSON(
		cfg.Connection,
		routing.ExchangeEzraDirect,
		routing.ActiveUserKey+"."+user.Name,
		routing.ActiveUserKey,
		msgbroker.SimpleQueueTransient,
		handlerActiveUsers(cfg),
	)

	return nil
}
