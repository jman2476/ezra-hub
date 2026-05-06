package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
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

	user, err := cfg.Client.LoginUser(loginData)
	if err != nil {
		return err
	}
	cfg.User = user
	printUser(user)

	return nil
}
