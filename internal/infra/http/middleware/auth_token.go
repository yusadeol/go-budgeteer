package middleware

import (
	"context"
	"errors"
	"github.com/yusadeol/go-budgeteer/internal/app/usecase"
	"github.com/yusadeol/go-budgeteer/internal/infra/adapter"
	"net/http"
	"strings"
)

type ContextKey string

const (
	ContextKeyUserId ContextKey = "userId"
)

type AuthToken struct {
}

func NewAuthToken() *AuthToken {
	return &AuthToken{}
}

func (m *AuthToken) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		token, err := m.parseToken(authorizationHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenParser := adapter.NewJWtParser()
		authToken := usecase.NewAuthToken(tokenParser)
		input := usecase.NewAuthTokenInput(token)
		var output *usecase.AuthTokenOutput
		output, err = authToken.Execute(input)
		if err != nil || output.Subject == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), ContextKeyUserId, output.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (*AuthToken) parseToken(authorizationHeader string) (string, error) {
	tokenPrefix := "Bearer "
	if !strings.HasPrefix(authorizationHeader, tokenPrefix) {
		return "", errors.New("invalid authorization header format")
	}
	token := strings.TrimPrefix(authorizationHeader, tokenPrefix)
	token = strings.TrimSpace(token)
	return token, nil
}
