package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	OwnerID   uuid.UUID `json:"owner_id"`
	Type      string    `json:"type"`
	OccursOn  time.Time `json:"occurs_on"`
	ExpiresAt time.Time `json:"expires_at"`
}

func mapEvent(e database.Event) Event {
	return Event{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Name:      e.Name,
		OwnerID:   e.OwnerID,
		Type:      e.Type,
		OccursOn:  e.OccursOn,
		ExpiresAt: e.ExpiresAt,
	}
}
