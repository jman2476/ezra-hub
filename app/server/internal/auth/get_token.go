package auth

import (
	"errors"
	"net/http"
	"strings"
)

var (
	malformedHeaderErr    = errors.New("Malformed token in header")
	missingBearerTokenErr = errors.New("Header does not contain token")
)

func GetBearerToken(headers http.Header) (string, error) {
	bearerStr := headers.Get("Authorization")
	if bearerStr == "" {
		return "", missingBearerTokenErr
	}

	token, ok := strings.CutPrefix(bearerStr, "Bearer ")
	if !ok {
		return "", malformedHeaderErr
	}

	return token, nil
}
