package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Method string

func (m Method) String() string {
	return string(m)
}

const (
	MethodGet     Method = "GET"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodDelete  Method = "DELETE"
	MethodPatch   Method = "PATCH"
	MethodOptions Method = "OPTIONS"
	MethodHead    Method = "HEAD"
	MethodConnect Method = "CONNECT"
	MethodTrace   Method = "TRACE"
)

type Server interface {
	Register(method Method, url string, callback func(w http.ResponseWriter, r *http.Request))
	Listen(port string) error
}

type ChiServer struct {
	router *chi.Mux
	server *http.Server
}

func NewChiServer() *ChiServer {
	server := &ChiServer{router: chi.NewRouter()}

	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.RequestID)
	server.router.Use(middleware.RealIP)
	server.router.Use(middleware.Timeout(60 * time.Second))

	return server
}

func (s *ChiServer) Register(method Method, url string, callback func(w http.ResponseWriter, r *http.Request)) {
	s.router.Method(method.String(), url, http.HandlerFunc(callback))
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
