package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Subscriptions struct {
	Subs map[string]int `json:"subscriptions"`
}

func (c *Client) SetSubscriptions(subs map[string]int) ([]string, error) {
	url := c.baseURL + "/api/users"

	var subStruct Subscriptions
	subStruct.Subs = subs

	subData, err := json.Marshal(subStruct)
	if err != nil {
		return []string{}, fmt.Errorf("Data marshalling error: %w", err)
	}
	body := bytes.NewReader(subData)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return []string{}, fmt.Errorf("Error creating request: %w", err)
	}
	req.Header = c.MakeAuthHeader(c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []string{}, fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatusError[[]string](res, "Error setting user subscriptions")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, fmt.Errorf("Error reading response body: %w", err)
	}

	var newSubscriptions []string
	err = json.Unmarshal(data, &newSubscriptions)
	if err != nil {
		return []string{}, fmt.Errorf("Error marshalling response body: %w", err)
	}

	return newSubscriptions, nil
}
