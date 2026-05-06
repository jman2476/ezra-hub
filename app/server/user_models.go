package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}

func mapUser(user database.User) User {
	return User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdateAt:    user.UpdateAt,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
}
