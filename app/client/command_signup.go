package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandSignUp(cfg *config) error {
	fmt.Println("Signup new user\n\r")
	var newUser apicaller.NewUser

	cfg.Term.SetPrompt("Name: ")
	name, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newUser.Name = name

	cfg.Term.SetPrompt("Email: ")
	email, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newUser.Email = email

	cfg.Term.SetPrompt("Phone number: ")
	phonenumber, err := cfg.Term.ReadLine()
	if err != nil {
		return err
	}
	newUser.PhoneNumber = phonenumber

	user, err := cfg.Client.NewUser(newUser)
	if err != nil {
		return err
	}
	fmt.Println("\r\n", user)
	printUser(user)

	return nil
}

func printUser(user apicaller.User) {
	fmt.Printf("\rName: %s\n", user.Name)
	fmt.Printf("\rCreated At: %v\n", user.CreatedAt)
	fmt.Printf("\rUpdated At: %v\n", user.UpdatedAt)
	fmt.Printf("\rEmail: %s\n", user.Email)
	fmt.Printf("\rPhone Number: %s\n", user.PhoneNumber)
}
