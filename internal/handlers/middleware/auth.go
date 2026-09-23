package middleware

import (
	"net/http"
	"context"
	"errors"
	"strings"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
)

// Auth Middleware
type AuthMiddleware struct {
	firebase *firebase.Client
}

func NewAuthMiddleware (firebase *firebase.Client) *AuthMiddleware {
	return &AuthMiddleware{
		firebase: firebase,
	}
}


func (m *AuthMiddleware) Authenticate(
	next http.Handler,
) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authClient, err := m.firebase.GetAuthClient(r.Context())
		
		if err != nil {
			http.Error(w, "Failed to fetch firebase authentication Client", http.StatusInternalServerError)
			return
		}

		token, err := m.extractToken(r)

		decodedToken, err := authClient.VerifyIDToken(
			r.Context(),
			token,
		)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Put user information into context
		ctx := context.WithValue(
			r.Context(),
			"user",
			decodedToken,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header must use Bearer scheme")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))

	if token == "" {
		return "", errors.New("token missing")
	}

	return token, nil
}
