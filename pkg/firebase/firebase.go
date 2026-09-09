package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)


type Client struct {
	app *firebase.App
}

func NewClient (credentialsFile string) (*Client, error) {
	opt := option.WithAuthCredentialsFile(
		option.ServiceAccount,
		credentialsFile,
		)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize firebase: %w", err)
	}
	return &Client{app: app}, nil
}

// Verify user Token
func (c *Client) VerifyToken(ctx context.Context, idToken string) (*auth.Token, error) {
	authClient, err := c.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Auth Client: %w", err)
	}
	//get token
	token, err := authClient.VerifyIDToken(ctx, idToken)
    if err != nil {
        return nil, fmt.Errorf("failed to verify token: %w", err)
    }

    return token, nil
}

//get Auth Client
func (c *Client) GetAuthClient(ctx context.Context) (*auth.Client, error) {
    return c.app.Auth(ctx)
}

