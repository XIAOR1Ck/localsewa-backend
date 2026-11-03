package handlers

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
)

type Handlers struct {
	svc             *service.SVC
	jwtSvc          *auth.JWTService
	AuthHandler     *AuthHandler
	ProviderHandler *ProviderHandler
}

func InitHandlers(svc *service.SVC, jwtSvc *auth.JWTService) *Handlers {
	return &Handlers{
		svc:    svc,
		jwtSvc: jwtSvc,
		AuthHandler: &AuthHandler{
			authService: svc.AuthService,
			jwtSvc:      jwtSvc,
		},
		ProviderHandler: &ProviderHandler{
			svc: svc.ProviderService,
		},
	}
}
