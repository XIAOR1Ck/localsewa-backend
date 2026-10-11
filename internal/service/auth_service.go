package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/XIAOR1Ck/localsewa-backend/internal/db"
	"github.com/XIAOR1Ck/localsewa-backend/internal/domain"
	"github.com/XIAOR1Ck/localsewa-backend/pkg/firebase"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5"
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

func (s *AuthService) hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func (s *AuthService) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

/* THIS SECTION CONTAINS FIREBASE AUTHENTICATION FUNCTIONS

//Verify user token for authentication
func (s *AuthService) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
    return s.firebaseClient.VerifyToken(ctx, idToken)
}

// Create user with Email and Password
func (s *AuthService) CreateUserWEmail(
	ctx context.Context,
	data domain.UserRegisterModel,
) (string, error) {

	client, err := s.firebaseClient.GetAuthClient(ctx)
	if err != nil {
		return "", fmt.Errorf("error fetching auth client: %w", err)
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
		return "", fmt.Errorf("error creating user: %w", err)
	}

	// Set initial custom claims
	claims := map[string]interface{}{
		"role":               "customer",
	}

	err = client.SetCustomUserClaims(ctx, u.UID, claims)
	if err != nil {
		return "", fmt.Errorf("error setting custom claims: %w", err)
	}
// Fetch the updated UserRecord with custom claims
u, err = client.GetUser(ctx, u.UID)
if err != nil {
	return "", fmt.Errorf("error fetching updated user: %w", err)
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
		return "", fmt.Errorf("Failed to add user to database: %w", err)
	}
token, err := client.CustomToken(ctx, u.UID)
if err != nil {
	return "", fmt.Errorf("error minting custom token: %w\n", err)
}

	fmt.Printf("User: %v", user)

return token, nil

}

// Create Service Provider with email and password
func (s *AuthService) CreateProviderWEmail(
	ctx context.Context,
	data domain.UserRegisterModel,
) (string, error) {

	client, err := s.firebaseClient.GetAuthClient(ctx)
	if err != nil {
		return "", fmt.Errorf("error fetching auth client: %w", err)
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
		return "", fmt.Errorf("error creating user: %w", err)
	}

	// Set initial custom claims
	claims := map[string]interface{}{
		"role":               "customer",
	}

	err = client.SetCustomUserClaims(ctx, u.UID, claims)
	if err != nil {
		return "", fmt.Errorf("error setting custom claims: %w", err)
	}
// Fetch the updated UserRecord with custom claims
u, err = client.GetUser(ctx, u.UID)
if err != nil {
	return "", fmt.Errorf("error fetching updated user: %w", err)
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
		return "", fmt.Errorf("Failed to add user to database: %w", err)
	}
token, err := client.CustomToken(ctx, u.UID)
if err != nil {
	return "", fmt.Errorf("error minting custom token: %w\n", err)
}

	fmt.Printf("User: %v", user)

return token, nil

}

*/

// Register Customer in postgres
func (s *AuthService) UserRegisterWEmail(
	ctx context.Context,
	data domain.UserRegisterModel,
) (domain.User, error) {
	_, err := s.repo.GetUserByEmail(ctx, data.Email)
	if err == nil {
    // User already exists
    return domain.User{}, errors.New("user with this email already exists")
} else if err != pgx.ErrNoRows {
    // Some other database error
    return domain.User{}, fmt.Errorf("failed to check existing user: %w", err)
}

	passHash, err := s.hashPassword(data.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("Failed to hash password: %w", err)
	}
	userData := db.CreateUserParams{
		Name: data.Name,
		Email: data.Email,
		Password: passHash,
		Phone: data.Phone,
		Role: db.UserRoleCustomer,
	}
	// Add user to database
	newUser, err := s.repo.CreateUser(ctx, userData)
	if err != nil {
		return domain.User{}, fmt.Errorf("Failed to create user: %w", err)
	}
	return domain.DbUserToModel(newUser), nil
}

// Register Provider in postgres
func (s *AuthService) ProviderRegisterWEmail(
	ctx context.Context,
	data domain.ProviderRegisterModel,
) (domain.Provider, error) {
	_, err := s.repo.GetUserByEmail(ctx, data.Email)
	if err == nil {
    // User already exists
    return domain.Provider{}, errors.New("user with this email already exists")
} else if err != pgx.ErrNoRows {
    // Some other database error
    return domain.Provider{}, fmt.Errorf("failed to check existing user: %w", err)
}

	passHash, err := s.hashPassword(data.Password)
	if err != nil {
		return domain.Provider{}, fmt.Errorf("Failed to hash password: %w", err)
	}
	userData := db.CreateProviderParams{
		Name: data.Name,
		Email: data.Email,
		Password: passHash,
		Phone: data.Phone,
	}
	// Add user to database
	newUser, err := s.repo.CreateProvider(ctx, userData)
	if err != nil {
		return domain.Provider{}, fmt.Errorf("Failed to create user: %w", err)
	}
	return domain.DbProviderToModel(newUser), nil
}


