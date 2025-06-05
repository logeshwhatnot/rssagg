package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	apiKeyString := headers.Get("Authorization")
	if apiKeyString == "" {
		return "", errors.New("no authorization value found")
	}
	apiKeyVal := strings.Split(apiKeyString, " ")
	if len(apiKeyVal) != 2 {
		return "", errors.New("malformed auth header")
	}
	if apiKeyVal[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return apiKeyVal[1], nil
}
