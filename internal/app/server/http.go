package server

import "net/http"

type Method string

func (m Method) String() string {
	return string(m)
}

const (
	MethodGet  Method = "GET"
	MethodPost Method = "POST"
)

type Http interface {
	Register(method Method, url string, callback func(w http.ResponseWriter, r *http.Request))
	Group(callback func(RouterGroup))
	Listen(port string) error
}

type RouterGroup interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	Register(method Method, url string, callback func(w http.ResponseWriter, r *http.Request))
}
