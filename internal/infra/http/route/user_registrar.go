package route

import (
	"database/sql"
	"github.com/yusadeol/go-budgeteer/internal/app/server"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/handler"
	"github.com/yusadeol/go-budgeteer/internal/infra/http/middleware"
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

	middlewareAuthToken := middleware.NewAuthToken()

	serverHttp.Register(server.MethodPost, "/users", userHandler.Store)
	serverHttp.Register(server.MethodPost, "/users/auth", userHandler.Auth)
	serverHttp.Group(func(group server.RouterGroup) {
		group.Use(middlewareAuthToken.Handle)

		group.Register(server.MethodGet, "/users", userHandler.Show)
	})
}
