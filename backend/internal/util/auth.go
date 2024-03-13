package util

import (
	"errors"
	"strings"
)

const authorizationTypeBearer = "Bearer"

func ValidateBearerHeader(authHeader string) (string, error) {
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return "", errors.New("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		return "", errors.New("unsupported authorization type")
	}

	return fields[1], nil
}
