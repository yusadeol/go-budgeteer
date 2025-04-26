package usecase

import (
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"testing"
)

type mockUserRepository struct {
	users []*entity.User
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Save(user *entity.User) error {
	m.users = append(m.users, user)
	return nil
}

type mockPasswordHasher struct{}

func newMockPasswordHasher() *mockPasswordHasher {
	return &mockPasswordHasher{}
}

func (m mockPasswordHasher) Hash(password string) (string, error) {
	return "mocked_hash", nil
}

func (m mockPasswordHasher) Verify(password, hash string) (bool, error) {
	return true, nil
}

func TestCreateUser_Execute(t *testing.T) {
	t.Run("should be able to create a new user", func(t *testing.T) {
		t.Parallel()

		userRepository := newMockUserRepository()
		passwordHasher := newMockPasswordHasher()
		useCase := NewCreateUser(userRepository, passwordHasher)

		input := NewCreateUserInput("Anthony Stark", "tony@marvel.com", "pepper123")

		output, err := useCase.Execute(input)
		if err != nil {
			t.Fatalf("unexpected error executing use case: %v", err)
		}

		if len(userRepository.users) == 0 {
			t.Fatal("expected at least one user to be saved, but got none")
		}

		lastUser := userRepository.users[len(userRepository.users)-1]
		if lastUser.Id != output.UserId {
			t.Fatalf("expected user ID %s, but got %s", output.UserId, lastUser.Id)
		}
	})
}
