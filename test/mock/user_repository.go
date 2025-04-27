package mock

import (
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"github.com/yusadeol/go-budgeteer/internal/infra/repository"
)

type UserRepository struct {
	Users []*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (m *UserRepository) Save(user *entity.User) error {
	m.Users = append(m.Users, user)
	return nil
}

func (m *UserRepository) GetByEmail(email string) (*entity.User, error) {
	for _, user := range m.Users {
		if user.Email != email {
			continue
		}
		return user, nil
	}

	return nil, repository.ErrUserNotFound
}
