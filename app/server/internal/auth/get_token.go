package auth

import (
	"errors"
	"log"
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
		log.Printf("Bearer string: %s\nToken: %s", bearerStr, token)
		return "", malformedHeaderErr
	}

	return token, nil
}
