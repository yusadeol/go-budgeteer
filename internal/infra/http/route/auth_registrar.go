package route

import (
	"github.com/yusadeol/go-budgeteer/internal/infra/http"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
)

type AuthRegistrar struct{}

func NewAuthRegistrar() *AuthRegistrar {
	return &AuthRegistrar{}
}

func (a *AuthRegistrar) Execute(server http.Server) {
	authHandler := handler.NewAuth()

	server.Register(http.MethodGet, "/auth/token", authHandler.GenerateToken)
}
