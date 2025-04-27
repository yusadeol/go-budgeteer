package usecase

import (
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
	"github.com/yusadeol/go-budgeteer/test/mock"
	"testing"
)

func TestAuthUser_Execute(t *testing.T) {
	t.Run("should be able to authenticate a user", func(t *testing.T) {
		t.Parallel()

		passwordHasher := mock.NewPasswordHasher()
		password := vo.NewPassword(passwordHasher)

		err := password.Parse("pepper123")
		if err != nil {
			t.Fatalf("failed to parse password: %v", err)
		}

		userRepository := mock.NewUserRepository()

		user := entity.NewUser("Anthony Stark", "tony@marvel.com", password)

		err = userRepository.Save(user)
		if err != nil {
			t.Fatalf("failed to save user: %v", err)
		}

		useCase := NewAuthUser(userRepository)

		input := NewAuthUserInput("tony@marvel.com", "pepper123")

		output, err := useCase.Execute(input)
		if err != nil {
			t.Fatalf("failed to authenticate user: %v", err)
		}

		if output.Id != user.Id {
			t.Fatalf("expected user ID %s, got %s", user.Id, output.Id)
		}
	})
}
