package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

type User struct {
	ID          uuid.UUID               `json:"id"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	Name        string                  `json:"name"`
	PhoneNumber string                  `json:"phone_number"`
	Email       string                  `json:"email"`
	JWT         string                  `json:"jwt"`
	Refresh     string                  `json:"refresh_token"`
	Subs        []database.Subscription `json:"subs"`
	Address     string                  `json:"address"`
}

func mapUser(user database.User, jwt, refresh string) User {
	if jwt != "" || refresh != "" {
		return User{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			JWT:         jwt,
			Refresh:     refresh,
			Subs:        user.Subs,
			Address:     user.Address,
		}
	}
	return User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Subs:        user.Subs,
		Address:     user.Address,
	}
}
