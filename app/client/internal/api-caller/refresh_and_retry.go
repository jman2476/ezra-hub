package apicaller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var (
	errRefreshedTooSoon = errors.New("Refreshing too soon, refresh token likely invalid")
	errBadRefreshToken  = errors.New("Refresh token expired or invalid, user must sign in again")
)

func (c *Client) Refresh() (status int, errVal error) {
	fmt.Println("Refreshing JWT\r")
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
		return handleStatusError[int](res, "Error refreshing JWT")
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
	fmt.Printf("\rRetrying create %s\r\n", newData.GetLogName())
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

func RetryEventRespond(c *Client, eventID uuid.UUID, available bool) error {
	fmt.Println("\rRetrying respond to event")
	if time.Since(c.lastRefresh) <= time.Minute {
		return errRefreshedTooSoon
	}

	status, err := c.Refresh()
	if err != nil {
		if status == 200 {
			return err
		}

		return errBadRefreshToken
	}

	return c.RespondtoEvent(eventID, available, true)
}

func RetryGetUserEvents(c *Client, categories []string) ([]Event, error) {
	fmt.Println("\rRetrying get user events")
	if time.Since(c.lastRefresh) <= time.Minute {
		return []Event{}, errRefreshedTooSoon
	}

	status, err := c.Refresh()
	if err != nil {
		if status == 200 {
			return []Event{}, err
		}
		return []Event{}, errBadRefreshToken
	}

	return c.GetUserEvents(categories, true)
}

func RetryGetUserCreatedEvents(c *Client) ([]Event, error) {
	fmt.Println("\rRetrying get user events")
	if time.Since(c.lastRefresh) <= time.Minute {
		return []Event{}, errRefreshedTooSoon
	}

	status, err := c.Refresh()
	if err != nil {
		if status == 200 {
			return []Event{}, err
		}
		return []Event{}, errBadRefreshToken
	}

	return c.GetUserCreatedEvents(true)
}

func RetrySubscribe(c *Client, subs map[string]int) ([]string, error) {
	fmt.Println("\rRetrying set subscriptions")
	if time.Since(c.lastRefresh) <= time.Minute {
		return []string{}, errRefreshedTooSoon
	}

	status, err := c.Refresh()
	if err != nil {
		if status == 200 {
			return []string{}, err
		}
		return []string{}, errBadRefreshToken
	}

	return c.SetSubscriptions(subs, true)
}
