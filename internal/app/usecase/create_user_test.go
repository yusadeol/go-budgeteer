package usecase

import (
	"github.com/yusadeol/go-budgeteer/test/mock"
	"testing"
)

func TestCreateUser_Execute(t *testing.T) {
	t.Run("should be able to create a new user", func(t *testing.T) {
		t.Parallel()

		userRepository := mock.NewUserRepository()
		passwordHasher := mock.NewPasswordHasher()
		useCase := NewCreateUser(userRepository, passwordHasher)

		input := NewCreateUserInput("Anthony Stark", "tony@marvel.com", "pepper123")

		output, err := useCase.Execute(input)
		if err != nil {
			t.Fatalf("unexpected error executing use case: %v", err)
		}

		if len(userRepository.Users) == 0 {
			t.Fatal("expected at least one user to be saved, but got none")
		}

		lastUser := userRepository.Users[len(userRepository.Users)-1]
		if lastUser.Id != output.Id {
			t.Fatalf("expected user ID %s, but got %s", output.Id, lastUser.Id)
		}
	})
}
