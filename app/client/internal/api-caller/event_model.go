package apicaller

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type NewEvent struct {
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	OccursOn      time.Time `json:"occurs_on"`
	ExpiresAt     time.Time `json:"expires_at"`
	MinVolunteers int32     `json:"min_volunteer"`
	MaxVolunteers int32     `json:"max_volunteer"`
	Description   string    `json:"description"`
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
	ID            uuid.UUID     `json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Name          string        `json:"name"`
	OwnerID       uuid.UUID     `json:"owner_id"`
	Category      string        `json:"category"`
	OccursOn      time.Time     `json:"occurs_on"`
	ExpiresAt     time.Time     `json:"expires_at"`
	MinVolunteers sql.NullInt32 `json:"min_volunteer"`
	MaxVolunteers sql.NullInt32 `json:"max_volunteer"`
	Description   string        `json:"description"`
	Respondants   []uuid.UUID   `json:"respondants"`
}
