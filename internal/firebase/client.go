package firebase

import (
	"context"
	"encoding/base64"
	"errors"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type ClientInterface interface {
	GetClient(ctx context.Context) *auth.Client
	GenerateCustomToken(ctx context.Context, identifier string, methodType AuthType) (string, error)
}

type Client struct {
	Credential string
	AppAuth    *auth.Client
}

func (c *Client) GetClient(ctx context.Context) *auth.Client {
	if c.AppAuth != nil {
		return c.AppAuth
	}

	decodedFirebaseCredential, err := base64.StdEncoding.DecodeString(c.Credential)
	if err != nil {
		log.Fatal("error:", err)
	}

	opt := option.WithCredentialsJSON(decodedFirebaseCredential)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}

	c.AppAuth = authClient

	return authClient
}

func (c *Client) GenerateCustomToken(ctx context.Context, identifier string, methodType AuthType) (string, error) {
	client := c.GetClient(ctx)

	var user *auth.UserRecord
	var err error

	if methodType == PhoneNumber {
		user, err = client.GetUserByPhoneNumber(ctx, identifier)
	} else if methodType == Email {
		user, err = client.GetUserByEmail(ctx, identifier)
	} else {
		return "", errors.New("invalid method type")
	}
	if err != nil {
		log.Printf("failed to find user: %v", err)
		return "", errors.New("failed to find user")
	}

	customToken, err := client.CustomToken(ctx, user.UID)
	if err != nil {
		log.Printf("failed to create custom token for user: %v", err)
		return "", errors.New("failed to generate token")
	}

	return customToken, nil
}
