package main

import (
	"fmt"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func commandSignUp(cfg *config) error {
	fmt.Println("Signup new user\n\r")
	var newUser apicaller.NewUser
	var user apicaller.User

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

	for {
		password, err := cfg.Term.ReadPassword("Set password: ")
		if err != nil {
			return err
		}

		retype, err := cfg.Term.ReadPassword("Retype password: ")
		if err != nil {
			return err
		}

		if password == retype {
			newUser.Password = password
			break
		}
	}

	user, err = cfg.Client.NewUser(newUser)
	if err != nil {
		return err
	}
	fmt.Println("\r\n", user)
	printUser(user)

	cfg.User = user
	//cheese

	return nil
}

func printUser(user apicaller.User) {
	fmt.Printf("\rName: %s\n", user.Name)
	fmt.Printf("\rCreated At: %v\n", user.CreatedAt)
	fmt.Printf("\rUpdated At: %v\n", user.UpdatedAt)
	fmt.Printf("\rEmail: %s\n", user.Email)
	fmt.Printf("\rPhone Number: %s\n", user.PhoneNumber)
	if user.Token != "" {
		fmt.Printf("\rJWT: %s\n", user.Token)
	}
	if user.Refresh != "" {
		fmt.Printf("\rRefresh: %s\n", user.Refresh)
	}
}
