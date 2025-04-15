package adapter

import (
	"errors"
	"github.com/yusadeol/go-budgeteer/internal/app"
	"testing"
)

func TestJWTGenerator_Execute(t *testing.T) {
	generator := NewJWTGenerator()

	t.Run("should generate token with valid key and subject", func(t *testing.T) {
		key := "thisisavalidkey123"
		subject := "user@example.com"

		token, err := generator.Execute(key, subject)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if token == "" {
			t.Fatal("expected token, got empty string")
		}
	})

	t.Run("should return error for short key", func(t *testing.T) {
		key := "shortkey"
		subject := "user@example.com"

		_, err := generator.Execute(key, subject)
		if !errors.Is(err, app.ErrInvalidKey) {
			t.Fatalf("expected ErrInvalidKey, got %v", err)
		}
	})

	t.Run("should return error for short subject", func(t *testing.T) {
		key := "thisisavalidkey123"
		subject := "a"

		_, err := generator.Execute(key, subject)
		if !errors.Is(err, app.ErrInvalidSubject) {
			t.Fatalf("expected ErrInvalidSubject, got %v", err)
		}
	})
}
