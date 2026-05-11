package apicaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	errRefreshedTooSoon = errors.New("Refreshing too soon, refresh token likely invalid")
	errBadRefreshToken  = errors.New("Refresh token expired or invalid, user must sign in again")
)

func (c *Client) Refresh() (status int, errVal error) {
	url := c.baseURL + "/api/refresh"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return 0, fmt.Errorf("Error creating refresh token request: %w", err)
	}

	req.Header = c.MakeAuthHeader(c.refresh)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Refresh JWT error: %w", err)
	}
	defer res.Body.Close()

	status = res.StatusCode

	if status != 200 {
		errResp := struct {
			Error string `json:"error"`
		}{}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			errVal = fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
			return
		}

		err = json.Unmarshal(data, &errResp)
		if err != nil {
			errVal = fmt.Errorf("Response code: %s, Error unmarshalling response body: %w", res.Status, err)
			return
		}

		errVal = fmt.Errorf("Error refreshing JWT: %s, %s", res.Status, errResp.Error)
		return
	}

	new_jwt := struct {
		Token string `json:"jwt_token"`
	}{}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		errVal = fmt.Errorf("Error reading response body: %w", err)
		return
	}

	err = json.Unmarshal(data, &new_jwt)
	if err != nil {
		errVal = fmt.Errorf("Error unmarshalling response body: %w", err)
		return
	}

	c.token = new_jwt.Token
	c.lastRefresh = time.Now()
	return
}

func RetryCreateNew[R any, NR NewResource[R]](c *Client, newData NR) (R, error) {
	var nilRes R
	if time.Since(c.lastRefresh) <= time.Minute {
		return nilRes, errRefreshedTooSoon
	}

	status, err := c.Refresh()
	if err != nil {
		if status == 200 {
			return nilRes, err
		}

		return nilRes, errBadRefreshToken
	}

	return CreateNewResource[R](c, newData, true)
}
