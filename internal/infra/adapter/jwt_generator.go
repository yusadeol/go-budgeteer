package adapter

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTGenerator struct{}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (a *JWTGenerator) Execute(key, subject string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "my-auth-server",
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return s, nil
}
