package route

import (
	"database/sql"
	"github.com/yusadeol/go-budgeteer/internal/app/server"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
)

type UserRegistrar struct {
	dbConnection *sql.DB
}

func NewUserRegistrar(dbConnection *sql.DB) *UserRegistrar {
	return &UserRegistrar{dbConnection: dbConnection}
}

func (r *UserRegistrar) Execute(serverHttp server.Http) {
	passwordHasher := adapter.NewBcryptHasher(10)
	userHandler := handler.NewUser(r.dbConnection, passwordHasher)

	serverHttp.Register(server.MethodPost, "/users", userHandler.Store)
	serverHttp.Register(server.MethodPost, "/users/auth", userHandler.Auth)
}
