package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandUpdateUser(cfg *config) error {
	var userUpdate apicaller.NewUser

	fmt.Println("\rType in the fields you want to update, leave the others blank\r")

	cfg.Term.SetPrompt("Name: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if name != "" {
		userUpdate.Name = name
	}

	cfg.Term.SetPrompt("Email: ")
	email, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if email != "" {
		userUpdate.Email = email
	}

	cfg.Term.SetPrompt("Phone number: ")
	phonenumber, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if phonenumber != "" {
		userUpdate.PhoneNumber = phonenumber
	}

	cfg.Client.UpdateUser(userUpdate)

	return nil
}
