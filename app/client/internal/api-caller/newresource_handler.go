package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewResource[Resource any] interface {
	GetLogName() string
	GetEndpointURL(*Client) string
}

func CreateNewResource[R any, NR NewResource[R]](c *Client, newData NR) (R, error) {
	url := newData.GetEndpointURL(c)
	var resource R

	resourceData, err := json.Marshal(newData)
	if err != nil {
		return resource, fmt.Errorf("Data marshalling error: %w", err)
	}
	body := bytes.NewReader(resourceData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return resource, fmt.Errorf("Error creating request: %w", err)
	}

	if c.token != "" {
		header := c.MakeAuthHeader(c.token)
		req.Header = header
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return resource, fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		errResp := struct {
			Error string `json:"error"`
		}{}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return resource, fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return resource, fmt.Errorf("Response code: %s, Error unmarshalling response body: %w", res.Status, err)
		}

		return resource, fmt.Errorf("Error creating %s: %s, %s", newData.GetLogName(), res.Status, errResp.Error)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return resource, fmt.Errorf("Error reading response body: %w", err)
	}

	err = json.Unmarshal(data, &resource)
	if err != nil {
		return resource, fmt.Errorf("Error unmarshalling response body: %w", err)
	}

	return resource, nil
}
