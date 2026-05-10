package apicaller

import (
	"time"

	"github.com/google/uuid"
)

type NewEvent struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	OccursOn  time.Time `json:"occurs_on"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (e NewEvent) GetLogName() string {
	return "new event"
}

func (e NewEvent) GetEndpointURL(c *Client) string {
	return c.baseURL + "/api/events"
}

func (e NewEvent) NewEmptyStruct() interface{} {
	return Event{}
}

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
