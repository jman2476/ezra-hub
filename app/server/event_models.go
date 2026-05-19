package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/ezra-hub/app/server/internal/database"
)

type Genre string

const (
	GenreRide      Genre = "ride"
	GenreShopping  Genre = "shopping"
	GenreCheckIn   Genre = "check-in"
	GenreMeal      Genre = "meal"
	GenreGathering Genre = "gathering"
	GenreOther     Genre = "other"
)

type Event struct {
	ID            uuid.UUID     `json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Name          string        `json:"name"`
	OwnerID       uuid.UUID     `json:"owner_id"`
	Category      Genre         `json:"category"`
	OccursOn      time.Time     `json:"occurs_on"`
	ExpiresAt     time.Time     `json:"expires_at"`
	MinVolunteers sql.NullInt32 `json:"min_volunteer"`
	MaxVolunteers sql.NullInt32 `json:"max_volunteer"`
	Description   string        `json:"description"`
	Respondants   []uuid.UUID   `json:"respondants"`
}

type EventwName struct {
	Event   `json:"event"`
	Creator string `json:"creator"`
}

func mapEvent(e database.Event) Event {
	return Event{
		ID:            e.ID,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		Name:          e.Name,
		OwnerID:       e.OwnerID,
		Category:      Genre(e.Category),
		OccursOn:      e.OccursOn,
		ExpiresAt:     e.ExpiresAt,
		MinVolunteers: e.MinVolunteers,
		MaxVolunteers: e.MaxVolunteers,
		Description:   e.Description,
		Respondants:   e.Respondants,
	}
}

func mapEventwName(e database.CreateEventRow) EventwName {
	return EventwName{
		Event: Event{
			ID:            e.ID,
			CreatedAt:     e.CreatedAt,
			UpdatedAt:     e.UpdatedAt,
			Name:          e.Name,
			OwnerID:       e.OwnerID,
			Category:      Genre(e.Category),
			OccursOn:      e.OccursOn,
			ExpiresAt:     e.ExpiresAt,
			MinVolunteers: e.MinVolunteers,
			MaxVolunteers: e.MaxVolunteers,
			Description:   e.Description,
			Respondants:   e.Respondants,
		},
		Creator: e.CreatorName,
	}
}

func validateGenre(s string) Genre {
	switch s {
	case "ride":
		return GenreRide
	case "shopping":
		return GenreShopping
	case "check-in":
		return GenreCheckIn
	case "meal":
		return GenreMeal
	case "gathering":
		return GenreGathering
	default:
		return GenreOther
	}
}
