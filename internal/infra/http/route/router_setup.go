package route

import (
	"github.com/yusadeol/go-budgeteer/internal/app/server"
)

type RouterSetup struct {
	server     server.Http
	registrars []RouteRegistrar
}

type RouteRegistrar interface {
	Execute(server server.Http)
}

func NewRouterSetup(server server.Http) *RouterSetup {
	return &RouterSetup{server: server}
}

func (r *RouterSetup) Register(registrars ...RouteRegistrar) {
	r.registrars = append(r.registrars, registrars...)
}

func (r *RouterSetup) Apply() {
	for _, registrator := range r.registrars {
		registrator.Execute(r.server)
	}
}
