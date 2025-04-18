package handler

import "net/http"

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (auth *Auth) GenerateToken(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}
