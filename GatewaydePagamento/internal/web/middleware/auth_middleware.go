package middleware

import (
	"net/http"

	"github.com/vivianezzt/go-gateway.git/internal/domain"
	"github.com/vivianezzt/go-gateway.git/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "API Key is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByAPIKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, "Account not found", http.StatusUnauthorized)
			return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		next.ServeHTTP(w, r)
	})
}

