package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/yusadeol/go-budgeteer/internal/app/usecase"
	"github.com/yusadeol/go-budgeteer/internal/domain/vo"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"github.com/yusadeol/go-budgeteer/internal/infra/repository"
	"net/http"
	"os"
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

	userRepository := repository.NewUserDatabase(h.dbConnection, h.passwordHasher)
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

func (h *User) Auth(w http.ResponseWriter, r *http.Request) {
	var authUserInput usecase.AuthUserInput

	userRepository := repository.NewUserDatabase(h.dbConnection, h.passwordHasher)
	authUser := usecase.NewAuthUser(userRepository)

	err := json.NewDecoder(r.Body).Decode(&authUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	authUserOutput, err := authUser.Execute(&authUserInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokenGenerator := adapter.NewJWTGenerator()
	generateToken := usecase.NewGenerateToken(tokenGenerator)

	tokenKey := os.Getenv("TOKEN_KEY")
	if tokenKey == "" {
		http.Error(w, "invalid token key", http.StatusBadRequest)
	}

	generateTokenInput := usecase.NewGenerateTokenInput(tokenKey, authUserOutput.Id)

	var generateTokenOutput *usecase.GenerateTokenOutput
	generateTokenOutput, err = generateToken.Execute(generateTokenInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&generateTokenOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
