package adapter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yusadeol/go-budgeteer/internal/app/server"
)

type ChiServer struct {
	router *chi.Mux
	server *http.Server
}

func NewChiServer() *ChiServer {
	chiServer := &ChiServer{router: chi.NewRouter()}

	chiServer.router.Use(middleware.Logger)
	chiServer.router.Use(middleware.Recoverer)
	chiServer.router.Use(middleware.RequestID)
	chiServer.router.Use(middleware.RealIP)
	chiServer.router.Use(middleware.Timeout(60 * time.Second))

	return chiServer
}

func (s *ChiServer) Register(method server.Method, url string, callback func(w http.ResponseWriter, r *http.Request)) {
	s.router.Method(method.String(), url, http.HandlerFunc(callback))
}

func (s *ChiServer) Group(callback func(server.RouterGroup)) {
	chiRouterGroup := NewChiRouterGroup(s.router)
	callback(chiRouterGroup)

	chiRouterGroup.router.Group(func(r chi.Router) {
		r.Use(chiRouterGroup.middlewares...)
		
		for _, route := range chiRouterGroup.routes {
			r.Method(route.method.String(), route.url, http.HandlerFunc(route.callback))
		}
	})
}

func (s *ChiServer) Listen(port string) error {
	addr := fmt.Sprintf(":%s", port)

	s.server = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.server.ListenAndServe()
}

type ChiRouterGroup struct {
	router      *chi.Mux
	middlewares []func(http.Handler) http.Handler
	routes      []*Route
}

func NewChiRouterGroup(router *chi.Mux) *ChiRouterGroup {
	return &ChiRouterGroup{
		router:      router,
		middlewares: []func(http.Handler) http.Handler{},
		routes:      []*Route{},
	}
}

func (rg *ChiRouterGroup) Use(middlewares ...func(http.Handler) http.Handler) {
	rg.middlewares = append(rg.middlewares, middlewares...)
}

func (rg *ChiRouterGroup) Register(method server.Method, url string, callback func(w http.ResponseWriter, r *http.Request)) {
	rg.routes = append(rg.routes, NewRoute(method, url, callback))
}

type Route struct {
	method   server.Method
	url      string
	callback func(w http.ResponseWriter, r *http.Request)
}

func NewRoute(method server.Method, url string, callback func(w http.ResponseWriter, r *http.Request)) *Route {
	return &Route{
		method:   method,
		url:      url,
		callback: callback,
	}
}
