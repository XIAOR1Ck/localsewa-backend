package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
	"github.com/golang-jwt/jwt/v5"
)


type JWTCustomClaims struct {
	Id 			string `json:"id"`
	Email		string `json:"email"`
	Role 		string `json:"role"`
	jwt.RegisteredClaims
}

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


// Generate Token With JWT
func (m *AuthMiddleware) GenerateJWT(claims JWTCustomClaims) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
// Set standard claims if not already set
	if claims.ExpiresAt == nil {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	}
	if claims.IssuedAt == nil {
		claims.IssuedAt = jwt.NewNumericDate(time.Now())
	}
	if claims.NotBefore == nil {
		claims.NotBefore = jwt.NewNumericDate(time.Now())
	}

	// Create token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Failed to sign token %w", err)
	}

	return tokenString, nil
}

