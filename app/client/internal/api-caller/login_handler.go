package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	return user, nil
}
