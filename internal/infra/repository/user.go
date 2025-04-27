package repository

import (
	"database/sql"
	"errors"
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Save(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
}

type UserDatabase struct {
	dbConnection *sql.DB
}

func NewUserDatabase(dbConnection *sql.DB) *UserDatabase {
	return &UserDatabase{dbConnection: dbConnection}
}

func (u *UserDatabase) Save(user *entity.User) error {
	_, err := u.dbConnection.Exec(
		"INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)",
		user.Id,
		user.Name,
		user.Email,
		user.Password.Value,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDatabase) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.dbConnection.QueryRow(
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?",
		email,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
