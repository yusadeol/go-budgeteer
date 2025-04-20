package route

import (
	"github.com/yusadeol/go-budgeteer/internal/infra/http"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
)

type UserRegistrar struct{}

func NewUserRegistrar() *UserRegistrar {
	return &UserRegistrar{}
}

func (a *UserRegistrar) Execute(server http.Server) {
	userHandler := handler.NewUser()

	server.Register(http.MethodPost, "/user", userHandler.Store)
}
