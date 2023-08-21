package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get API key from http header
// Authorize: ApiKey {key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorize")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("wrong auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("wrong first part of auth header")
	}

	return vals[1], nil
}
