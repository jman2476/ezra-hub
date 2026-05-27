package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetUserCreatedEvents(retry bool) ([]Event, error) {
	url := c.baseURL + "/api/events/users"

	if c.token == "" {
		return []Event{}, fmt.Errorf("Cannot get user events: no user signed in")
	}

	body := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return []Event{}, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header = c.MakeAuthHeader(c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []Event{}, fmt.Errorf("Error getting events created by current user")
	}
	res.Body.Close()

	if res.StatusCode == 204 {
		return []Event{}, fmt.Errorf("User owns no current events")
	}

	if !retry && res.StatusCode == 401 {
		return RetryGetUserCreatedEvents(c)
	}

	if res.StatusCode != 200 {
		return handleStatusError[[]Event](res, "Error getting user owned events")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []Event{}, fmt.Errorf("Error reading response body: %w", err)
	}
	var events []Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		return []Event{}, fmt.Errorf("Error unmarshalling response body: %w", err)
	}

	return events, nil
}
