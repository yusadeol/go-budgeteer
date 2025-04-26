package repository

import (
	"database/sql"
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
)

type User interface {
	Save(user *entity.User) error
}

type UserDatabase struct {
	dbConnection *sql.DB
}

func NewUserDatabase(dbConnection *sql.DB) *UserDatabase {
	return &UserDatabase{dbConnection: dbConnection}
}

func (u *UserDatabase) Save(user *entity.User) error {
	stmt, err := u.dbConnection.Prepare(`
		INSERT INTO users(id, name, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Password.Value, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
