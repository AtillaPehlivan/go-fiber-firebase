package auth

import (
	"context"
	"firebase.google.com/go/auth"
	"mocerize-api/pkg/firebase"
)

var (
	firebaseAuthClient *auth.Client
)

func Setup(ctx context.Context) error {

	authClient, err := firebase.GetFirebaseApp().Auth(ctx)

	if err != nil {
		return err
	}

	firebaseAuthClient = authClient
	return nil
}

func Client() *auth.Client {
	return firebaseAuthClient
}
