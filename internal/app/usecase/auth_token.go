package usecase

import (
	"errors"
)

type AuthToken struct {
	tokenParser TokenParser
}

type TokenParser interface {
	Execute(token string) (map[string]any, error)
}

func NewAuthToken(tokenParser TokenParser) *AuthToken {
	return &AuthToken{tokenParser: tokenParser}
}

func (u *AuthToken) Execute(input *AuthTokenInput) (*AuthTokenOutput, error) {
	claims, err := u.tokenParser.Execute(input.Token)
	if err != nil {
		return nil, err
	}
	sub, isString := claims["sub"].(string)
	if !isString {
		return nil, errors.New("sub claim is missing or not a string")
	}
	return NewAuthTokenOutput(sub), nil
}

type AuthTokenInput struct {
	Token string
}

func NewAuthTokenInput(token string) *AuthTokenInput {
	return &AuthTokenInput{Token: token}
}

type AuthTokenOutput struct {
	Subject string
}

func NewAuthTokenOutput(subject string) *AuthTokenOutput {
	return &AuthTokenOutput{Subject: subject}
}
