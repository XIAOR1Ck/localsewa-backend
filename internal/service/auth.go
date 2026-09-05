package service

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
)

type AuthService interface {
    VerifyToken(ctx context.Context, idToken string) (*auth.Token, error)
		CreateUserWEmail(ctx context.Context, name string, email string, phone string, password string)(*auth.UserRecord, error)
}

type FirebaseAuthService struct {
    firebaseClient *firebase.Client
}

func NewAuthService(firebaseClient *firebase.Client) *FirebaseAuthService {
    return &FirebaseAuthService{
        firebaseClient: firebaseClient,
    }
}

//Verify user token for authentication
func (s *FirebaseAuthService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
    return s.firebaseClient.VerifyToken(ctx, idToken)
}

// Create user with Email and Password
func (s *FirebaseAuthService) CreateUserWEmail(ctx context.Context, data domain.UserRegister) (*auth.UserRecord, error) {
	client, err := s.firebaseClient.GetAuthClient(ctx)
	
	if err != nil {
		return nil, fmt.Errorf("Error fetching auth client: %w", err)
	}

	// Params to create user on
	params := (&auth.UserToCreate{}).
	Email(data.Email).
	EmailVerified(false).
	PhoneNumber(data.Phone).
	Password(data.Password).
	DisplayName(data.Name).
	Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return u, nil
}
