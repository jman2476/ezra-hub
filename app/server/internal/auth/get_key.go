package auth

import (
	"errors"
	"net/http"
	"strings"
)

var (
	MissingAPIKeyErr   = errors.New("missing API key")
	MalformedAPIKeyErr = errors.New("malformed API key")
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKeyStr := headers.Get("Authorization")
	if apiKeyStr == "" {
		return "", MissingAPIKeyErr
	}

	key, ok := strings.CutPrefix(apiKeyStr, "ApiKey ")
	if !ok {
		return "", MalformedAPIKeyErr
	}

	return key, nil
}
