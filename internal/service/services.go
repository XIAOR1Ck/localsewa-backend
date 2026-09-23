package service

import (
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
)

type SVC struct {
	repo *db.Queries
	firebase  *firebase.Client
	AuthService *AuthService
	// Add other services here
}

func InitServices(repo *db.Queries, firebase *firebase.Client) *SVC {
	return &SVC{
		repo: repo,
		firebase: firebase,
		AuthService: NewAuthService(firebase, repo),
	}
}
