package apicaller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handleStatusError[T any](res *http.Response, finalErr string) (T, error) {
	errResp := struct {
		Error string `json:"error"`
	}{}
	var empty T

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return empty, fmt.Errorf("Response code: %s, Error reading response body: %w", res.Status, err)
	}
	err = json.Unmarshal(data, &errResp)
	if err != nil {
		return empty, fmt.Errorf("Response code: %s, Error unmarshalling response body: %w", res.Status, err)
	}

	return empty, fmt.Errorf("%s: %s, %s", finalErr, res.Status, errResp.Error)
}
