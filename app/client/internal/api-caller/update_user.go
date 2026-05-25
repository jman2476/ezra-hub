package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) UpdateUser(nu UserUpdate) (User, error) {
	url := c.baseURL + "/api/users"

	updateData, err := json.Marshal(nu)
	if err != nil {
		return User{}, fmt.Errorf("Data marshaling error: %w", err)
	}
	body := bytes.NewReader(updateData)

	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return User{}, fmt.Errorf("Error creating request: %w", err)
	}

	header := c.MakeAuthHeader(c.token)
	req.Header = header

	res, err := c.httpClient.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatusError[User](res, "Error updating user data")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return User{}, fmt.Errorf("Error reading response body: %w", err)
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return User{}, fmt.Errorf("Error unmarshaling response body: %w", err)
	}

	return user, nil

	// v := reflect.ValueOf(nu)
	// fields := v.Fields()

	// for data, field := range fields {
	// 	if field.Interface() != "" {
	// 		fmt.Printf("\rNew %s: %s\n", data.Name, field.Interface())
	// 	}
	// }

}
