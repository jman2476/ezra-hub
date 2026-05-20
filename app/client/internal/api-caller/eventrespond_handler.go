package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) RespondtoEvent(eventID uuid.UUID, available, retry bool) error {
	url := c.baseURL + "/api/events/" + eventID.String()

	var going = struct {
		Available bool `json:"available"`
	}{
		Available: available,
	}

	respondData, err := json.Marshal(going)
	if err != nil {
		return fmt.Errorf("Data marshalling error: %w", err)
	}
	body := bytes.NewReader(respondData)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return fmt.Errorf("Error creating request: %w", err)
	}
	req.Header = c.MakeAuthHeader(c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	// Insert retry
	if !retry && res.StatusCode == 401 {
		return RetryEventRespond(c, eventID, available)
	}

	if res.StatusCode != 204 {
		_, err = handleStatusError[any](res, "Error responding to event")
		return err
		// errResp := struct {
		// 	Error string `json:"error"`
		// }{}

		// data, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	return fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		// }

		// err = json.Unmarshal(data, &errResp)
		// if err != nil {
		// 	return fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		// }

		// return fmt.Errorf("Error logging in user: %s, %s", res.Status, errResp.Error)
	}

	fmt.Println("\rResponse successfully sent to server")

	return nil
}
