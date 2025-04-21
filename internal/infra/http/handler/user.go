package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/yusadeol/go-budgeteer/internal/app/usecase"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
	"github.com/yusadeol/go-budgeteer/internal/infra/repository"
	"net/http"
)

type User struct {
	dbConnection   *sql.DB
	passwordHasher vo.PasswordHasher
}

func NewUser(dbConnection *sql.DB, passwordHasher vo.PasswordHasher) *User {
	return &User{dbConnection: dbConnection, passwordHasher: passwordHasher}
}

func (h *User) Store(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	var output *usecase.CreateUserOutput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userRepository := repository.NewUser(h.dbConnection)
	createUser := usecase.NewCreateUser(userRepository, h.passwordHasher)

	output, err = createUser.Execute(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
