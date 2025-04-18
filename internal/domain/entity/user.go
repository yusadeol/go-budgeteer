package entity

import "github.com/google/uuid"

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	return &User{
		Id:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}
