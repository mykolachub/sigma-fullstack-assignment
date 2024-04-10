package util

import (
	"sigma-user/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(id, role, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.JWTClaimsId:   id,
		config.JWTClaimsRole: role,
		config.JWTClaimsExp:  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ParseAndValidateJWTToken(accessToken, secret string) (*jwt.Token, config.ServiceCode, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, config.SvcInvalidToken, config.SvcInvalidToken.ToError()
	}

	return token, config.SvcEmptyMsg, nil
}
