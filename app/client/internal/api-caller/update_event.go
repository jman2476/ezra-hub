package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) UpdateEvent(eventInfo EventUpdate, eventID uuid.UUID) (Event, error) {
	url := c.baseURL + "/api/events/" + eventID.String()

	updateData, err := json.Marshal(eventInfo)
	if err != nil {
		return Event{}, fmt.Errorf("Data marshaling error: %w", err)
	}
	body := bytes.NewReader(updateData)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return Event{}, fmt.Errorf("Error creating request: %w", err)
	}

	header := c.MakeAuthHeader(c.token)
	req.Header = header

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Event{}, fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatusError[Event](res, "Error updating event data")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Event{}, fmt.Errorf("Error reading response body: %w", err)
	}

	var event Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		return Event{}, fmt.Errorf("Error unmarshaling response body: %w", err)
	}

	return event, nil
}
