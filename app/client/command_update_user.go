package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandUpdateUser(cfg *config) error {
	var userUpdate apicaller.UserUpdate

	fmt.Println("\rType in the fields you want to update, leave the others blank\r")

	cfg.Term.SetPrompt("Name: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if name != "" {
		userUpdate.Name = name
	} else {
		userUpdate.Name = cfg.User.Name
	}

	cfg.Term.SetPrompt("Email: ")
	email, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if email != "" {
		userUpdate.Email = email
	} else {
		userUpdate.Email = cfg.User.Email
	}

	cfg.Term.SetPrompt("Phone number: ")
	phonenumber, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if phonenumber != "" {
		userUpdate.PhoneNumber = phonenumber
	} else {
		userUpdate.PhoneNumber = cfg.User.PhoneNumber
	}

	cfg.Term.SetPrompt("Address: ")
	address, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	if address != "" {
		userUpdate.Address = address
	} else {
		userUpdate.Address = cfg.User.Address
	}

	user, err := cfg.Client.UpdateUser(userUpdate)
	if err != nil {
		return fmt.Errorf("Update user client error: %w", err)
	}

	cfg.User = user

	return nil
}
