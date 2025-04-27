package usecase

import (
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
	"github.com/yusadeol/go-budgeteer/internal/infra/repository"
)

type CreateUser struct {
	userRepository repository.User
	passwordHasher vo.PasswordHasher
}

func NewCreateUser(userRepository repository.User, passwordHasher vo.PasswordHasher) *CreateUser {
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
	Id string `json:"id"`
}

func NewCreateUserOutput(id string) *CreateUserOutput {
	return &CreateUserOutput{
		Id: id,
	}
}
