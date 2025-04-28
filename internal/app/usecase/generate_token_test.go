package usecase

import (
	"github.com/yusadeol/go-budgeteer/test/mock"
	"testing"
)

func TestGenerateToken_Execute(t *testing.T) {
	t.Run("should be able to generate an token", func(t *testing.T) {
		t.Parallel()

		tokenGenerator := mock.NewTokenGenerator()
		useCase := NewGenerateToken(tokenGenerator)
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
}
