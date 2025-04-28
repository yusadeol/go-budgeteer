package mock

type TokenGenerator struct{}

func NewTokenGenerator() *TokenGenerator {
	return &TokenGenerator{}
}

func (m *TokenGenerator) Execute(key, subject string) (string, error) {
	return "fake.jwt.token", nil
}
