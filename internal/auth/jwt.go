package auth

import (
	"time"
	"fmt"
	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	ID 			string `json:"id"`
	Email		string `json:"email"`
	Role 		string `json:"role"`
	jwt.RegisteredClaims
}

type JWTService struct {	
	cfg 		config.Config
}

func NewJWTService (
	cfg config.Config,
) *JWTService {
	return &JWTService{
	cfg: cfg,	
	}
}

// Generate Token With JWT
func (s *JWTService) Sign(ID, email, role string) (string, error) {
	secretKey := s.cfg.JWTSecret
	now := time.Now()
	claims := JWTCustomClaims{
		ID: ID,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: s.cfg.JWTIssuer,
			IssuedAt: jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.JWTTTL)),
		},
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
