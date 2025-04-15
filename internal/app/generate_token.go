package app

import "errors"

var (
	ErrInvalidKey     = errors.New("invalid key")
	ErrInvalidSubject = errors.New("invalid subject")
)

type GenerateToken struct {
	TokenGenerator TokenGenerator
}

type TokenGenerator interface {
	Execute(key, subject string) (string, error)
}

func NewGenerateToken(TokenGenerator TokenGenerator) *GenerateToken {
	return &GenerateToken{
		TokenGenerator: TokenGenerator,
	}
}

func (u *GenerateToken) Execute(input *GenerateTokenInput) (*GenerateTokenOutput, error) {
	t, err := u.TokenGenerator.Execute(input.Key, input.Subject)
	if err != nil {
		return nil, err
	}

	return NewGenerateTokenOutput(t), nil
}

type GenerateTokenInput struct {
	Key     string
	Subject string
}

func NewGenerateTokenInput(key, subject string) *GenerateTokenInput {
	return &GenerateTokenInput{
		Key:     key,
		Subject: subject,
	}
}

type GenerateTokenOutput struct {
	Token string
}

func NewGenerateTokenOutput(token string) *GenerateTokenOutput {
	return &GenerateTokenOutput{
		Token: token,
	}
}
