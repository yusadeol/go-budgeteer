package middleware

import (
	"fmt"
	"net/http"
)

type AuthToken struct {
}

func NewAuthToken() *AuthToken {
	return &AuthToken{}
}

func (*AuthToken) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		fmt.Println(token)

		next.ServeHTTP(w, r)
	})
}
