package route

import "github.com/yusadeol/go-budgeteer/internal/infra/http"

type RouterSetup struct {
	server     http.Server
	registrars []RouteRegistrar
}

type RouteRegistrar interface {
	Execute(server http.Server)
}

func NewRouterSetup(server http.Server) *RouterSetup {
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
