package entity

import (
	"github.com/google/uuid"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     string
	Password  *vo.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email string, password *vo.Password) *User {
	return &User{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
