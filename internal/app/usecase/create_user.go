package usecase

import (
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
)

type CreateUser struct {
	userRepository UserRepository
	passwordHasher vo.PasswordHasher
}

type UserRepository interface {
	Save(user *entity.User) error
}

func NewCreateUser(userRepository UserRepository, passwordHasher vo.PasswordHasher) *CreateUser {
	return &CreateUser{userRepository: userRepository, passwordHasher: passwordHasher}
}

func (u *CreateUser) Execute(input *CreateUserInput) (*CreateUserOutput, error) {
	password := vo.NewPassword(u.passwordHasher)
	err := password.Parse(input.Password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(input.Name, input.Email, password)

	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	output := NewCreateUserOutput(user.Id)
	return output, nil
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserInput(name, email, password string) *CreateUserInput {
	return &CreateUserInput{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

type CreateUserOutput struct {
	UserId string `json:"user_id"`
}

func NewCreateUserOutput(userId string) *CreateUserOutput {
	return &CreateUserOutput{
		UserId: userId,
	}
}
