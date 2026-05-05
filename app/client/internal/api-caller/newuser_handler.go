package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) NewUser(signupInfo newUser) (User, error) {
	url := c.baseURL + "/api/users"

	userData, err := json.Marshal(signupInfo)
	if err != nil {
		return User{}, err
	}
	body := bytes.NewReader(userData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return User{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
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

		return User{}, fmt.Errorf("Error creating new user: %s, %s", res.Status, errResp.Error)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
