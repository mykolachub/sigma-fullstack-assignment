package util

import (
	"sigma-test/config"
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

func ParseAndValidateJWTToken(accessToken, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, config.ErrInvalidToken
	}

	return token, nil
}
