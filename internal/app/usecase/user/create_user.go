package user

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

func (u *CreateUser) Execute(input Input) (*Output, error) {
	user := entity.NewUser(input.Name, input.Email, input.Password)

	err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	output := NewOutput(user.Id)
	return output, nil
}

type Input struct {
	Name     string
	Email    string
	Password string
}

func NewInput(name, email, password string) *Input {
	return &Input{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

type Output struct {
	UserId string
}

func NewOutput(userId string) *Output {
	return &Output{
		UserId: userId,
	}
}
