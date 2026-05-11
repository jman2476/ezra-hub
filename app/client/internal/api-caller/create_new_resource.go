package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewResource interface {
	GetLogName() string
	GetEndpointURL(*Client) string
	NewEmptyStruct() interface{}
}

type Resource interface {
	User | Event
}

func (c *Client) CreateNewResource(newData NewResource) (interface{}, error) {
	url := newData.GetEndpointURL(c)

	resourceData, err := json.Marshal(newData)
	if err != nil {
		return newData.NewEmptyStruct(), fmt.Errorf("Data marshalling error: %w", err)
	}
	body := bytes.NewReader(resourceData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return newData.NewEmptyStruct(), fmt.Errorf("Error creating request: %w", err)
	}

	header, ok := c.MakeAuthHeader()
	if ok {
		req.Header = header
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return newData.NewEmptyStruct(), fmt.Errorf("Error executing request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		errResp := struct {
			Error string `json:"error"`
		}{}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return newData.NewEmptyStruct(), fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return newData.NewEmptyStruct(), fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
		}

		return newData.NewEmptyStruct(), fmt.Errorf("Error creating %s: %s, %s", newData.GetLogName(), res.Status, errResp.Error)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return newData.NewEmptyStruct(), fmt.Errorf("Error reading response body: %w", err)
	}

	resource := newData.NewEmptyStruct()
	err = json.Unmarshal(data, &resource)
	if err != nil {
		return newData.NewEmptyStruct(), fmt.Errorf("Error unmarshalling response body: %w", err)
	}

	return resource, nil
}
