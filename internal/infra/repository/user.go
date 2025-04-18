package repository

import (
	"database/sql"
	"github.com/yusadeol/go-budgeteer/internal/domain/entity"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Save(user *entity.User) error {
	stmt, err := u.db.Prepare("INSERT INTO users(id, name, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
