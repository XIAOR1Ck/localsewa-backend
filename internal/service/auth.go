package service

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
)



type AuthService struct {
    firebaseClient *firebase.Client
		repo *db.Queries
}

func NewAuthService(firebaseClient *firebase.Client, repo *db.Queries) *AuthService {
    return &AuthService{
        firebaseClient: firebaseClient,
				repo: repo,
    }
}

//Verify user token for authentication
func (s *AuthService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
    return s.firebaseClient.VerifyToken(ctx, idToken)
}

// Create user with Email and Password
// Create user with Email and Password
func (s *AuthService) CreateUserWEmail(
	ctx context.Context,
	data domain.UserRegister,
) (db.User, error) {

	client, err := s.firebaseClient.GetAuthClient(ctx)
	if err != nil {
		return db.User{}, fmt.Errorf("error fetching auth client: %w", err)
	}

	// Create Firebase user
	params := (&auth.UserToCreate{}).
		Email(data.Email).
		EmailVerified(false).
		PhoneNumber(data.Phone).
		Password(data.Password).
		DisplayName(data.Name).
		Disabled(false)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return db.User{}, fmt.Errorf("error creating user: %w", err)
	}

	// Set initial custom claims
	claims := map[string]interface{}{
		"role":               "customer",
	}

	err = client.SetCustomUserClaims(ctx, u.UID, claims)
	if err != nil {
		return db.User{}, fmt.Errorf("error setting custom claims: %w", err)
	}
// Fetch the updated UserRecord with custom claims
u, err = client.GetUser(ctx, u.UID)
if err != nil {
	return db.User{}, fmt.Errorf("error fetching updated user: %w", err)
}

// Getting data ready
userData := &db.CreateUserParams{
		ID: u.UID,
		Name: u.DisplayName,
		Email: u.Email,
		Phone: u.PhoneNumber,
		Role: "customer",
	}

// Write to database

	user, err := s.repo.CreateUser(ctx, *userData)
	if err != nil {
		return db.User{}, fmt.Errorf("Failed to add user to database: %w", err)
	}

return user, nil

}
