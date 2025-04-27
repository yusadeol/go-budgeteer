package usecase

import (
	"errors"
	"github.com/yusadeol/go-budgeteer/internal/infra/repository"
)

var (
	ErrAuthUserInvalidPassword = errors.New("invalid password")
)

type AuthUser struct {
	userRepository repository.User
}

func NewAuthUser(userRepository repository.User) *AuthUser {
	return &AuthUser{userRepository: userRepository}
}

func (u *AuthUser) Execute(input *AuthUserInput) (*AuthUserOutput, error) {
	user, err := u.userRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	var isEqual bool
	isEqual, err = user.Password.Compare(input.Password)
	if !isEqual {
		return nil, ErrAuthUserInvalidPassword
	}

	return NewAuthUserOutput(user.Id), nil
}

type AuthUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthUserInput(email, password string) *AuthUserInput {
	return &AuthUserInput{
		Email:    email,
		Password: password,
	}
}

type AuthUserOutput struct {
	Id string `json:"id"`
}

func NewAuthUserOutput(id string) *AuthUserOutput {
	return &AuthUserOutput{
		Id: id,
	}
}
