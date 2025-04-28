package route

import (
	"database/sql"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"github.com/yusadeol/go-budgeteer/internal/infra/http"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
)

type UserRegistrar struct {
	dbConnection *sql.DB
}

func NewUserRegistrar(dbConnection *sql.DB) *UserRegistrar {
	return &UserRegistrar{dbConnection: dbConnection}
}

func (r *UserRegistrar) Execute(server http.Server) {
	passwordHasher := adapter.NewBcryptHasher(10)
	userHandler := handler.NewUser(r.dbConnection, passwordHasher)

	server.Register(http.MethodPost, "/users", userHandler.Store)
	server.Register(http.MethodPost, "/users/auth", userHandler.Auth)
}
