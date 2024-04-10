package util

import (
	"sigma-user/config"
	"strings"
)

const authorizationTypeBearer = "bearer"

func ValidateBearerHeader(authHeader string) (string, config.ServiceCode, error) {
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return "", config.SvcInvalidBearer, config.SvcInvalidBearer.ToError()
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		return "", config.SvcInvalidAuthType, config.SvcInvalidAuthType.ToError()
	}

	return fields[1], config.SvcEmptyMsg, nil
}
