package usecase

import "github.com/yusadeol/go-budgeteer/internal/domain/entity"

type CreateUser struct {
	userRepository UserRepository
}

type UserRepository interface {
	Save(user *entity.User) error
}

func NewCreateUser(userRepository UserRepository) *CreateUser {
	return &CreateUser{userRepository: userRepository}
}

func (u *CreateUser) Execute(input CreateUserInput) (*CreateUserOutput, error) {
	user := entity.NewUser(input.Name, input.Email, input.Password)

	err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	output := NewCreateUserOutput(user.Id)
	return output, nil
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func NewCreateUserInput(name, email, password string) *CreateUserInput {
	return &CreateUserInput{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

type CreateUserOutput struct {
	UserId string
}

func NewCreateUserOutput(userId string) *CreateUserOutput {
	return &CreateUserOutput{
		UserId: userId,
	}
}
