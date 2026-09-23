package handlers

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/service"
)

type Handlers struct {
	svc *service.SVC
	AuthHandler *AuthHandler
}

func InitHandlers (svc *service.SVC) *Handlers {
	return &Handlers{
		svc: svc,
		AuthHandler: &AuthHandler{
			authService: svc.AuthService,
		},
	}
}

