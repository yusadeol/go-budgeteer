package repository

import (
	"database/sql"
	"errors"
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Save(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
}

type UserDatabase struct {
	dbConnection   *sql.DB
	passwordHasher vo.PasswordHasher
}

func NewUserDatabase(dbConnection *sql.DB, passwordHasher vo.PasswordHasher) *UserDatabase {
	return &UserDatabase{dbConnection: dbConnection, passwordHasher: passwordHasher}
}

func (r *UserDatabase) Save(user *entity.User) error {
	_, err := r.dbConnection.Exec(
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

func (r *UserDatabase) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	var password string

	err := r.dbConnection.QueryRow(
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?",
		email,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	voPassword := vo.NewPasswordFromHash(r.passwordHasher, password)
	user.Password = voPassword

	return &user, nil
}
