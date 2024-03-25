package util

import (
	"errors"
	"sigma-test/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	env       = config.ConfigEnv()
	jwtSecret = env.JWTSecret
)

func GenerateJWTToken(id, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}

func ParseAndValidateJWTToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, err
}
