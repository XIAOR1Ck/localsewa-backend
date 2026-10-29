package middleware

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	jwtsvc *auth.JWTService
}

func NewAuthMiddleware(svc *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwtsvc: svc}
}

const claimsKey = "auth.claims"

func (m *AuthMiddleware) Require() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		raw, found := strings.CutPrefix(header, "Bearer ")
		if !found || raw == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		claims, err := m.jwtsvc.Verify(raw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set(claimsKey, claims)
		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := ClaimsFrom(c)
		if !ok || claims.Role != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}

// Typed accessor so handlers never touch the string key.
func ClaimsFrom(c *gin.Context) (*auth.JWTCustomClaims, bool) {
	v, ok := c.Get(claimsKey)
	if !ok {
		return nil, false
	}
	cl, ok := v.(*auth.JWTCustomClaims)
	return cl, ok
}
