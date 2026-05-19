package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) LoginUser(loginInfo UserLogin) (User, error) {
	url := c.baseURL + "/api/login"

	loginData, err := json.Marshal(loginInfo)
	if err != nil {
		return User{}, fmt.Errorf("Data marshalling error: %w", err)
	}
	body := bytes.NewReader(loginData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return User{}, fmt.Errorf("Error creating request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		errResp := struct {
			Error string `json:"error"`
		}{}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return User{}, fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}

		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return User{}, fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}

		return User{}, fmt.Errorf("Error logging in user: %s, %s", res.Status, errResp.Error)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return User{}, fmt.Errorf("Error reading response body: %w", err)
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return User{}, fmt.Errorf("Error unmarshalling response body: %w", err)
	}

	c.token = user.Token
	c.refresh = user.Refresh
	c.lastRefresh = time.Now()

	return user, nil
}

func (c *Client) GetUserEvents(categories []string) ([]Event, error) {
	url := c.baseURL + "/api/events?"

	for _, c := range categories {
		url += "type=" + c + "&"
	}

	body := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return []Event{}, fmt.Errorf("Error creating request: %w", err)
	}

	if c.token == "" {
		return []Event{}, fmt.Errorf("Error: need token to make this request")
	}
	req.Header = c.MakeAuthHeader(c.token)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []Event{}, fmt.Errorf("Error getting user events: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 204 {
		return []Event{}, nil
	}

	if res.StatusCode != 200 {
		errResp := struct {
			Error string `json:"error"`
		}{}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return []Event{}, fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}
		fmt.Printf("\rjson data: %v\n", data)
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return []Event{}, fmt.Errorf("Response code: %s, Error unmarshaling response body: %w", res.Status, err)
		}

		return []Event{}, fmt.Errorf("Error getting user events: %s, %s", res.Status, errResp.Error)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []Event{}, fmt.Errorf("Error reading response body: %w", err)
	}

	var eventList []Event
	err = json.Unmarshal(data, &eventList)
	if err != nil {
		return []Event{}, fmt.Errorf("Error unmarshalling response body: %w", err)
	}

	return eventList, nil
}
