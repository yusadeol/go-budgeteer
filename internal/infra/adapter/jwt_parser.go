package adapter

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type JWtParser struct{}

func NewJWtParser() *JWtParser {
	return &JWtParser{}
}

func (a *JWtParser) Execute(jwtToken string) (map[string]any, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("TOKEN_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	claims, isMapClaims := token.Claims.(jwt.MapClaims)
	if !isMapClaims {
		return nil, errors.New("invalid token claims type")
	}
	return claims, nil
}
