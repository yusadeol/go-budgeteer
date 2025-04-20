package handler

import (
	"encoding/json"
	"github.com/yusadeol/go-budgeteer/internal/app/usecase"
	"net/http"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (h *User) Store(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
