package usecase

import (
	"github.com/yusadeol/go-budgeteer/test/mock"
	"testing"
)

func TestAuthToken_Execute(t *testing.T) {
	const expectedSubject = "uuid"
	const mockToken = "mock-token"
	t.Run("should be able to authenticate a token", func(t *testing.T) {
		t.Parallel()
		mockTokenParser := mock.NewTokenParser()
		authToken := NewAuthToken(mockTokenParser)
		input := NewAuthTokenInput(mockToken)
		output, err := authToken.Execute(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if output.Subject != expectedSubject {
			t.Fatalf("expected subject %q, got %q", expectedSubject, output.Subject)
		}
	})
}
