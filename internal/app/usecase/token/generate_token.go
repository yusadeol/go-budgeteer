package token

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

func (u *GenerateToken) Execute(input *Input) (*Output, error) {
	t, err := u.TokenGenerator.Execute(input.Key, input.Subject)
	if err != nil {
		return nil, err
	}

	return NewOutput(t), nil
}

type Input struct {
	Key     string
	Subject string
}

func NewInput(key, subject string) *Input {
	return &Input{
		Key:     key,
		Subject: subject,
	}
}

type Output struct {
	Token string
}

func NewOutput(token string) *Output {
	return &Output{
		Token: token,
	}
}
