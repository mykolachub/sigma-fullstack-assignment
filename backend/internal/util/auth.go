package util

import (
	"sigma-test/config"
	"strings"
)

const authorizationTypeBearer = "bearer"

func ValidateBearerHeader(authHeader string) (string, error) {
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return "", config.ErrInvalidBearer
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		return "", config.ErrInvalidAuthType
	}

	return fields[1], nil
}
