package usecase

import (
	"errors"
	"testing"
)

type mockTokenGenerator struct{}

func newMockTokenGenerator() *mockTokenGenerator {
	return &mockTokenGenerator{}
}

func (m *mockTokenGenerator) Execute(key, subject string) (string, error) {
	if key != "test-key" {
		return "", ErrGenerateTokenInvalidKey
	}

	if subject != "user@example.com" {
		return "", ErrGenerateTokenInvalidSubject
	}

	return "fake.jwt.token", nil
}

func TestGenerateToken_Execute(t *testing.T) {
	t.Run("should be able to generate an token", func(t *testing.T) {
		t.Parallel()

		mock := newMockTokenGenerator()
		useCase := NewGenerateToken(mock)
		input := NewGenerateTokenInput("test-key", "user@example.com")

		expectedToken := "fake.jwt.token"

		output, err := useCase.Execute(input)
		if err != nil {
			t.Fatalf("expected token %q, got %+v", expectedToken, err)
		}

		if output.Token != expectedToken {
			t.Errorf("expected token %q, got %q", expectedToken, output.Token)
		}
	})

	t.Run("should return an error when key or subject is invalid", func(t *testing.T) {
		t.Parallel()

		mock := newMockTokenGenerator()
		useCase := NewGenerateToken(mock)

		input := NewGenerateTokenInput("invalid", "user@example.com")

		_, err := useCase.Execute(input)
		if !errors.Is(err, ErrGenerateTokenInvalidKey) {
			t.Errorf("expected error %+q, got %+v", ErrGenerateTokenInvalidKey, err)
		}

		input = NewGenerateTokenInput("test-key", "invalid")

		_, err = useCase.Execute(input)
		if !errors.Is(err, ErrGenerateTokenInvalidSubject) {
			t.Errorf("expected error %+q, got %+v", ErrGenerateTokenInvalidSubject, err)
		}
	})
}
