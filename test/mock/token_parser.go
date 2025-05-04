package mock

type TokenParser struct{}

func NewTokenParser() *TokenParser {
	return &TokenParser{}
}

func (m TokenParser) Execute(token string) (map[string]any, error) {
	return map[string]any{
		"sub": "uuid",
	}, nil
}
