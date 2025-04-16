package adapter

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/yusadeol/go-budgeteer/internal/app/usecase"
	"time"
)

type JWTGenerator struct{}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (a *JWTGenerator) Execute(key, subject string) (string, error) {
	if len(key) < 16 {
		return "", usecase.ErrInvalidKey
	}

	if len(subject) < 2 {
		return "", usecase.ErrInvalidSubject
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "my-auth-server",
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return s, nil
}
